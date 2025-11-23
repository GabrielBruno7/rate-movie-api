echo "🔄 Rodando migrate:fresh ..."

export $(grep -v '^#' .env | xargs)

DB_URL="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

echo "🧨 Derrubando banco..."
migrate -path=migrations -database "$DB_URL" drop -f

echo "🚀 Subindo migrations do zero..."
migrate -path=migrations -database "$DB_URL" up

echo "✨ Banco recriado do zero com sucesso (fresh)!"
