package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

type Ordem struct {
	Correlacao []map[string]interface{} `json:"correlacao"`
	Operacao   map[string]interface{}   `json:"operacao"`
	Cliente    map[string]interface{}   `json:"cliente"`
	ItemOrdem  []map[string]interface{} `json:"item-ordem"`
}

func ReadSQLFile(filePath string) string {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// ProcessTestCase verifica o tipo de conteúdo do arquivo e chama a função apropriada.
func ProcessTestFile(filePath string) (string, []string, string, error) {
	fileName := filepath.Base(filePath)                              // Pega apenas "arquivo.sql"
	testName := strings.TrimSuffix(fileName, filepath.Ext(fileName)) // Remove ".sql"

	sqlContent := ReadSQLFile(filePath) // Ler o conteúdo do arquivo SQL

	// Determinar qual processamento será aplicado
	switch {
	case regexp.MustCompile(`(?i)INSERT\s+INTO\s+PROV\s*_\s*Q`).MatchString(sqlContent):
		fmt.Println("🔍 Detected: PROV_Q - Chamando ProcessProvQ")
		return ProcessProvQ(testName, sqlContent)

	case regexp.MustCompile(`(?i)INSERT\s+INTO\s+SPS\s*_\s*SOLICITACAO`).MatchString(sqlContent):
		fmt.Println("🔍 Detected: SPS_SOLICITACAO - Chamando ProcessSolicitacao")
		return ProcessSolicitacao1(testName, sqlContent)

	case isJSON(sqlContent):
		fmt.Println("🔍 Detected: JSON structure - Chamando ProcessRestJson")
		return ProcessRestJson(testName, sqlContent)

	default:
		err := fmt.Errorf("❌ Nenhuma estrutura identificada no arquivo: %s", filePath)
		fmt.Println(err)
		return "", nil, "0", err
	}
}

// isJSON verifica se o conteúdo tem formato JSON válido (simples verificação).
func isJSON(content string) bool {
	content = strings.TrimSpace(content)
	return strings.HasPrefix(content, "{") && strings.HasSuffix(content, "}")
}

// Adiciona "<%TCID%>" como o primeiro valor de cada comando INSERT INTO PROV_Q.
func ProcessProvQ(testName, sqlContent string) (string, []string, string, error) {
	var processedLines []string

	// Expressão regular para encontrar comandos INSERT INTO
	insertRegex := regexp.MustCompile(`(?i)INSERT INTO\s+.+\s+VALUES\s*\((.+)\)`)

	// Separar o conteúdo em linhas
	lines := strings.Split(sqlContent, "\n")

	for _, line := range lines {
		// Verifica se a linha corresponde a um INSERT INTO
		if matches := insertRegex.FindStringSubmatch(line); matches != nil {
			values := matches[1] // Captura os valores dentro de VALUES (...)

			// Separar os valores e substituir o primeiro por "<%TCID%>"
			valueParts := strings.SplitN(values, ",", 2)
			if len(valueParts) > 1 {
				modifiedValues := "<%TCID%>, " + valueParts[1]
				modifiedLine := strings.Replace(line, values, modifiedValues, 1)
				processedLines = append(processedLines, modifiedLine)
			} else {
				// Caso não haja separação correta, apenas mantém a linha original
				processedLines = append(processedLines, line)
			}
		}
	}

	// Retornar o nome do teste e os inserts processados
	return testName, []string{strings.Join(processedLines, "\n")}, "2", nil
}

// Esta função extrai todos os comandos INSERT INTO SPS_SOLICITACAO (sem o ponto e vírgula final)
// do conteúdo SQL e os retorna como uma única string concatenada.
// O 'testName' de entrada é o nome base do arquivo e será retornado.
func ProcessSolicitacao1(testName, sqlContent string) (string, []string, string, error) {
	var extractedInserts []string // Para armazenar os scripts INSERT extraídos

	re := regexp.MustCompile(`(?is)INSERT\s+INTO\s+SPS_SOLICITACAO\s*\(.*?\)\s*VALUES\s*\(.*?\)`)
	matches := re.FindAllString(sqlContent, -1) // -1 para encontrar todas as ocorrências
	if len(matches) == 0 {
		return testName, nil, "3", fmt.Errorf("❌ Nenhum INSERT INTO SPS_SOLICITACAO encontrado no arquivo: %s", testName)
	}

	// Para cada INSERT encontrado, adicione-o diretamente à lista de inserts extraídos.
	for _, insertStatement := range matches {
		extractedInserts = append(extractedInserts, insertStatement)
	}

	// Concatena todos os INSERTs extraídos em uma única string, separados por uma quebra de linha.
	// Esta string concatenada será o 'processedSQL' retornado.
	processedSQL := strings.Join(extractedInserts, "\n")

	// Expressão regular para encontrar cada comando INSERT INTO SPS_SOLICITACAO.
	reqIDModifierRe := regexp.MustCompile(`(?is)(seq_seq_geral\.NEXTVAL,\s*'?0'?(?:,\s*)?)'.*?'(\s*,\s*20)`)
	modifiedProcessedSQL := reqIDModifierRe.ReplaceAllString(processedSQL, "$1<%TCID%>$2")
	insertSplitterRe := regexp.MustCompile(`(?is)INSERT\s+INTO\s+SPS_SOLICITACAO\s*\(.*?\)\s*VALUES\s*\(.*?\)`)
	inserts := insertSplitterRe.FindAllString(modifiedProcessedSQL, -1)

	// Retorna o 'testName' original, a string concatenada de INSERTs, o tipo "3", e nil (sem erro).
	return testName, inserts, "3", nil
}

// Altera os valores de req_id para <%TCID%> e separa os INSERT INTO SPS_SOLICITACAO em uma lista
func ProcessSolicitacao2(processedSQL string) []string {
	// Expressão regular para encontrar cada comando INSERT INTO SPS_SOLICITACAO.
	reqIDModifierRe := regexp.MustCompile(`(?is)(seq_seq_geral\.NEXTVAL,\s*'?0'?(?:,\s*)?)'.*?'(\s*,\s*20)`)
	modifiedProcessedSQL := reqIDModifierRe.ReplaceAllString(processedSQL, "$1<%TCID%>$2")
	insertSplitterRe := regexp.MustCompile(`(?is)INSERT\s+INTO\s+SPS_SOLICITACAO\s*\(.*?\)\s*VALUES\s*\(.*?\)`)
	inserts := insertSplitterRe.FindAllString(modifiedProcessedSQL, -1)

	return inserts
}

// Adiciona o valor <%TCID>% ao campo "id" dentro de "correlacao" que está dentro de "ordem"
func ProcessRestJson(testName, jsonContent string) (string, []string, string, error) {
	// Parse o JSON em uma estrutura genérica
	var jsonData interface{}
	err := json.Unmarshal([]byte(jsonContent), &jsonData)
	if err != nil {
		return "", nil, "", fmt.Errorf("erro ao parsear o JSON: %v", err)
	}

	// Função recursiva para percorrer o JSON e modificar o valor de "id" dentro de "correlacao", que está dentro de "ordem"
	var processMap func(map[string]interface{})
	processMap = func(mapData map[string]interface{}) {
		for key, value := range mapData {
			// Verificar se é "ordem" e contém "correlacao"
			if key == "ordem" {
				if ordemMap, ok := value.(map[string]interface{}); ok {
					if correlacaoArr, found := ordemMap["correlacao"].([]interface{}); found {
						for _, item := range correlacaoArr {
							if correlacaoMap, ok := item.(map[string]interface{}); ok {
								// Substituir qualquer valor de "id" por "<%TCID%>"
								if _, idExists := correlacaoMap["id"]; idExists {
									correlacaoMap["id"] = "<%TCID%>"
								}
							}
						}
					}
				}
			}
			// Se o valor for um mapa, chamar recursivamente
			if nestedMap, ok := value.(map[string]interface{}); ok {
				processMap(nestedMap)
			}
		}
	}

	// Aplicar a modificação se o JSON for um mapa
	if dataMap, ok := jsonData.(map[string]interface{}); ok {
		processMap(dataMap)

		// Criar um buffer para armazenar o JSON formatado sem HTML escaping
		var buf bytes.Buffer
		encoder := json.NewEncoder(&buf)
		encoder.SetEscapeHTML(false) // Evita que "<" e ">" sejam escapados
		encoder.SetIndent("", "  ")  // Mantém a formatação bonita

		// Codificar o JSON modificado no buffer
		err = encoder.Encode(dataMap)
		if err != nil {
			return "", nil, "", fmt.Errorf("erro ao criar item modificado: %v", err)
		}

		// Retornar o JSON formatado corretamente
		return testName, []string{buf.String()}, "1", nil
	}

	return "", nil, "", fmt.Errorf("... Estrutura inesperada de JSON")
}

func GeraDescription(script string) string {
	// Regex para IMSI
	regexIMSI := regexp.MustCompile(`IMSI=([^;]+);`)
	// Regex para MSISDN
	regexMSISDN := regexp.MustCompile(`MSISDN=([^;]+);`)
	// Nova Regex para HHUA, HLR ou HLREDA
	regexHLX := regexp.MustCompile(`(HHUA|HLREDA|HLR)=([^;]+);`)
	// Regex para VOLTE (case-insensitive) - \b garante que é a palavra completa
	regexVOLTE := regexp.MustCompile(`(?i)\bVOLTE\b`) // `(?i)` para case-insensitive
	// Regex para VPNSIX (case-insensitive)
	regexVPNSIX := regexp.MustCompile(`(?i)\bVPNSIX\b`)

	var description string
	var imsiExtracted string
	var msisdnExtracted string
	var hlxExtracted string // Variável para armazenar o valor de HHUA/HLR/HLREDA
	var volteFound bool     // Variável para indicar se VOLTE foi encontrado
	var vpnsixFound bool    // Variável para indicar se VPNSIX foi encontrado

	// --- Extração do IMSI ---
	matchesIMSI := regexIMSI.FindStringSubmatch(script)
	if len(matchesIMSI) > 1 {
		fullIMSI := matchesIMSI[1]
		// Garante que pegamos apenas os primeiros 7 caracteres, se existirem
		if len(fullIMSI) >= 7 {
			imsiExtracted = fullIMSI[:7]
		} else {
			imsiExtracted = fullIMSI // Se for menor que 7, pega tudo
		}
	} else {
		imsiExtracted = "" // Se não encontrar, deixa vazio
	}

	// --- Extração do MSISDN ---
	matchesMSISDN := regexMSISDN.FindStringSubmatch(script)
	if len(matchesMSISDN) > 1 {
		msisdnExtracted = matchesMSISDN[1]
	} else {
		msisdnExtracted = "" // Se não encontrar, deixa vazio
	}

	// --- Extração de HHUA/HLR/HLREDA ---
	matchesHLX := regexHLX.FindStringSubmatch(script)
	if len(matchesHLX) > 2 {
		hlxFieldName := matchesHLX[1] // Ex: "HHUA", "HLR", "HLREDA"
		hlxValue := matchesHLX[2]     // Ex: "ValorHHUA1"
		// Apenas alteramos a forma como hlxExtracted é formatado
		hlxExtracted = fmt.Sprintf("%s=%s", hlxFieldName, hlxValue)
	} else {
		hlxExtracted = ""
	}

	// --- Verificação VOLTE E VPNSIX ---
	// Usa FindString para verificar a existência, não precisa de submatches
	if regexVOLTE.FindString(script) != "" {
		volteFound = true
	}
	if regexVPNSIX.FindString(script) != "" {
		vpnsixFound = true
	}

	// --- Construção da Description ---
	// Concatenar apenas se os valores foram encontrados
	description = ""
	parts := []string{}
	if imsiExtracted != "" {
		parts = append(parts, "IMSI="+imsiExtracted)
	}
	if msisdnExtracted != "" {
		parts = append(parts, "MSISDN="+msisdnExtracted)
	}
	if hlxExtracted != "" { // Adiciona o campo HLX se for encontrado
		parts = append(parts, hlxExtracted)
	}
	if volteFound {
		parts = append(parts, "VOLTE") // Adiciona a string "VOLTE"
	}
	if vpnsixFound {
		parts = append(parts, "VPNSIX") // Adiciona a string "VPNSIX"
	}

	// Junta as partes com um separador, se houver mais de uma
	description = strings.Join(parts, ", ")

	return description
}

// Simular a valores processados no banco exibindo o SQL gerado
func SimulaInsert(idType, testName, description, processedSQL, userName, idProject, idSchemaSPS string) {
	fmt.Println("==== SQL QUE SERIA INSERIDO ====")
	fmt.Println("ID:", "Incrementado no banco")
	fmt.Println("VERSION:", 1)
	fmt.Println("ID Type:", idType)
	fmt.Println("Nome:", testName)
	fmt.Println("Description: ", description)
	fmt.Println("DT_CREATION: ", "SYSDATE")
	fmt.Println("DT_UPDATE: ", "SYSDATE")
	fmt.Println("Usuário:", userName)
	fmt.Println("Projeto:", idProject)
	fmt.Println("Schema SPS:", idSchemaSPS)
	fmt.Println("Script: \n", processedSQL)
	fmt.Println("================================")
}
