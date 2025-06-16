#!/bin/bash

# Define o sufixo a ser adicionado
sufixo="_HUA"

# Caminho da pasta (pode ser passado como argumento ou fixado)
pasta="./"  # Diretório atual, altere se quiser

# Percorre todos os arquivos .js na pasta
for arquivo in "$pasta"*.js; do
  # Verifica se existe algum arquivo correspondente
  if [ -e "$arquivo" ]; then
    # Extrai o nome base do arquivo (sem extensão e sem caminho)
    nome_base=$(basename "$arquivo" .js)

    # Monta o novo nome com o sufixo antes da extensão
    novo_nome="${nome_base}${sufixo}.js"

    # Renomeia o arquivo
    mv "$arquivo" "${pasta}${novo_nome}"

    echo "Arquivo '$arquivo' renomeado para '${novo_nome}'"
  fi
done
