# 🚀 Fullstack Go CRUD + Streamlit

Uma aplicação completa que une a robustez do **Go** no backend com a agilidade do **Streamlit** no frontend. O projeto utiliza **PostgreSQL** para persistência e é totalmente orquestrado via **Docker**, garantindo um ambiente de desenvolvimento isolado e profissional.

Este projeto reflete meus estudos em Engenharia de Software, aplicando conceitos de arquitetura de APIs, concorrência e integração entre diferentes tecnologias.

---

## 🛠️ Stack Tecnológica

### **Backend**
* **Linguagem:** [Go](https://go.dev/) (Golang)
* **Banco de Dados:** [PostgreSQL](https://www.postgresql.org/)
* **Monitoramento:** Health Checker nativo (Goroutines/Channels)

### **Frontend**
* **Interface:** [Streamlit](https://streamlit.io/) (Python framework)
* **Comunicação:** Consumo de API REST via Requests

### **Infraestrutura & Dev**
* **Containerização:** Docker & Docker Compose
* **Interface de DB:** DBeaver
* **Editor:** Neovim (LazyVim) no Pop!_OS (COSMIC)

---

## ✨ Funcionalidades

- [x] **CRUD Completo:** Interface amigável para gerenciar dados em tempo real.
- [x] **Backend Concorrente:** API em Go otimizada para múltiplas requisições.
- [x] **Dashboard Intuitivo:** Front-end limpo e funcional com Streamlit.
- [x] **Orquestração Docker:** API, DB e Front subindo com um único comando.

---

## 🚀 Como Executar o Projeto

### 1. Pré-requisitos
Certifique-se de ter o **Docker** e o **Docker Compose** instalados.

### 2. Clonando o Repositório

```
git clone https://github.com/alvarolucio/estoque-livros-py-go.git
cd estoque-livros-py-go
```

### 3. Subindo a Stack Completa

O Docker Compose vai configurar automaticamente a rede entre o Streamlit, a API em Go e o Postgres:

```
docker-compose up --build -d
```
### 👨‍💻 Autor

Desenvolvido por Álvaro – Estudante de Engenharia de Software e Backend Developer em formação.

Este projeto é open-source. Sinta-se à vontade para contribuir!
