#!/bin/bash

# Verifica que se pase una URL como argumento
if [ -z "$1" ]; then
  echo "Uso: ./extract_token.sh '<URL>'"
  exit 1
fi

# URL recibida como argumento
URL="$1"

# Extrae el access_token usando sed (compatible con macOS)
ACCESS_TOKEN=$(echo "$URL" | sed -n 's/.*access_token=\([^&]*\).*/\1/p')

# Verifica si se extrajo el token
if [ -z "$ACCESS_TOKEN" ]; then
  echo "No se encontró access_token en la URL."
  exit 1
fi

# Obtiene la ruta del directorio del script
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
OUTPUT_FILE="$SCRIPT_DIR/token.txt"

# Elimina el archivo si ya existe
[ -f "$OUTPUT_FILE" ] && rm "$OUTPUT_FILE"

# Guarda el token sin salto de línea
printf "%s" "$ACCESS_TOKEN" > "$OUTPUT_FILE"

echo "✅ Token extraído y guardado en:"
echo "$OUTPUT_FILE"

# Copia sin salto de línea al portapapeles
printf "%s" "$ACCESS_TOKEN" | pbcopy

# Verifica si se copió correctamente
CLIPBOARD_CONTENT=$(pbpaste)
if [ "$CLIPBOARD_CONTENT" = "$ACCESS_TOKEN" ]; then
  echo "✅ Token copiado al portapapeles."
else
  echo "⚠️  No se pudo copiar al portapapeles. Puedes copiarlo manualmente desde token.txt"
fi