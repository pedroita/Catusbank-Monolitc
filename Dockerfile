# Derivando da imagem oficial do MySQL
FROM mysql:5.7
# Adicionando os scripts SQL para serem executados na criação do banco
COPY ./database/ /docker-entrypoint-initdb.d/