package database

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	_ "github.com/godror/godror" // Importa o driver Oracle
)

// conecta no BD, usando a string de conexao + conta + senha informados
func Connect() (*sql.DB, error) {

	DbConnStr := "oracle://%s:%s@192.168.2.140:1521/FREEPDB1"
	DbUserName := "TESTRUNNER"
	DbPassword := "TESTRUNNER"

	// gera a string de conexao do oracle, com o nome do usuario e senha
	connStr := fmt.Sprintf(DbConnStr, DbUserName, DbPassword)
	fmt.Println("🌐 Conectando ao banco de dados", slog.String("connStr", connStr))

	// tenta abrir a conexao
	db, err := sql.Open("oracle", connStr)
	if err != nil {
		fmt.Println("❌ Erro conectando no BD: ", slog.Any("error", err))
		return nil, fmt.Errorf("❌ Erro conectando no BD: %w", err)
	}

	// testa a conexao
	dbCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = db.PingContext(dbCtx)
	if err != nil {
		fmt.Println("db.PingContext - erro fazendo ping no BD", slog.Any("error", err))
		return nil, fmt.Errorf("db.PingContext - erro fazendo ping no BD: %w", err)
	}
	return db, err
}

// InsertTestCase insere um novo registro na tabela test_case
func InsertTestCase(db *sql.DB, idType, testName, processedSQL, description, userName, idProject, idSchemaSPS string) error {

	// Inicia uma transação
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("Erro ao iniciar transação: %v", err)
	}

	query := `
INSERT INTO TESTRUNNER.TEST_CASE (VERSION, ID_TYPE, NAME, SCRIPT, DESCRIPTION, USER_NAME, ID_PROJECT, ID_SCHEMA_SPS) VALUES
(1, :1, :2, TO_CLOB(:3), :4, :5, :6, :7)`

	fmt.Println("... Executando query de INSERT...") // LOG PARA DEBUG

	// Executa o comando SQL
	res, err := tx.Exec(query, idType, testName, processedSQL, description, userName, idProject, idSchemaSPS)
	if err != nil {
		fmt.Printf("❌ Erro na execução do INSERT: %v\n", err) // CAPTURA O ERRO
		tx.Rollback()                                         // Reverte a transação
		return fmt.Errorf("Erro ao inserir no banco: %v", err)
	}

	// Confirma a transação
	err = tx.Commit()
	if err != nil {
		fmt.Printf("❌ Erro ao confirmar transação (commit): %v\n", err) // CAPTURA O ERRO
		return fmt.Errorf("Erro ao confirmar transação: %v", err)
	}

	// Verifica se a inserção realmente afetou alguma linha
	rowsAffected, _ := res.RowsAffected()
	fmt.Printf("... Linhas inseridas: %d\n", rowsAffected)

	if rowsAffected == 0 {
		fmt.Println("⚠ Nenhuma linha foi inserida. Pode ser um problema nos dados enviados.") // LOG PARA DEPURAÇÃO
		return fmt.Errorf("Nenhuma linha foi inserida, verifique os dados.")
	}

	fmt.Println("✅ Registro inserido com sucesso na tabela test_case!")
	return nil
}
