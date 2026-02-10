# Dockerfile
FROM python:3.12.3-slim

WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copia tudo para o container (src, front, app.py, etc)
COPY . .

# Expõe as portas padrão
EXPOSE 8000
EXPOSE 8501
