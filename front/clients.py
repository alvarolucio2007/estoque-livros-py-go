import requests
import os


API_URL = os.getenv("API_URL", "http://localhost:8000")


def _tratar_resposta(response):
    """Função auxiliar para validar o status code e tratar erros."""
    if response.status_code in [200, 201]:  # 200: OK, 201: Created
        return response.json()

    # Se chegou aqui, deu erro
    try:
        detalhe = response.json().get("detail", "Erro desconhecido")
    except ValueError:
        detalhe = response.text

    raise Exception(f"Falha na API ({response.status_code}): {detalhe}")


# --- GET (Ler) ---


def listar_livro():
    try:
        # Tenta bater na porta da API
        response = requests.get(
            f"{API_URL}/livros", timeout=10
        )  # Timeout é bom pra não travar o app pra sempre
        return _tratar_resposta(response)
    except requests.exceptions.ConnectionError:
        # Se ninguém atender a porta (API desligada/Rede caiu)
        raise Exception("Erro de Conexão: A API parece estar offline ou inalcançável.")
    except requests.exceptions.Timeout:
        raise Exception("Erro de Tempo: A API demorou demais para responder.")


def gerar_relatorio():
    try:
        response = requests.get(f"{API_URL}/livros/relatorio", timeout=10)
        return _tratar_resposta(response)
    except requests.exceptions.ConnectionError:
        raise Exception("Erro de Conexão: A API parece estar offline ou inalcançável.")
    except requests.exceptions.Timeout:
        raise Exception("Erro de Tempo: A API demorou demais para responder.")


def listar_id():
    try:
        response = requests.get(f"{API_URL}/livros/listar_id", timeout=10)
        return _tratar_resposta(response)
    except requests.exceptions.ConnectionError:
        raise Exception("Erro de Conexão: A API parece estar offline ou inalcançável.")
    except requests.exceptions.Timeout:
        raise Exception("Erro de Tempo: A API demorou demais para responder.")


def buscar_livro_codigo(livro_id: int):
    try:
        response = requests.get(f"{API_URL}/livros/{livro_id}")
        return _tratar_resposta(response)
    except requests.exceptions.ConnectionError:
        raise Exception("Erro de conexão ao tentar buscar o livro. ")


def buscar_livro_autor(autor: str):
    try:
        response = requests.get(f"{API_URL}/livros/autor/{autor}")
        return _tratar_resposta(response)
    except requests.exceptions.ConnectionError:
        raise Exception("Erro de conexão ao tentar buscar o livro. ")


def buscar_livro_titulo(titulo: str):
    try:
        response = requests.get(f"{API_URL}/livros/titulo/{titulo}")
        return _tratar_resposta(response)
    except requests.exceptions.ConnectionError:
        raise Exception("Erro de conexão ao tentar buscar o livro. ")


# --- POST (Criar) ---


def cadastrar_livro(livro_dados: dict):
    try:
        response = requests.post(f"{API_URL}/livros", json=livro_dados, timeout=10)
        return _tratar_resposta(response)
    except requests.exceptions.ConnectionError:
        raise Exception("Erro de conexão ao tentar cadastrar.")


# --- Patch (Editar) ---


def editar_titulo(id: int, novo_titulo: str):
    try:
        response = requests.patch(f"{API_URL}/livros/{id}/titulo?titulo={novo_titulo}")
        return _tratar_resposta(response)
    except requests.exceptions.ConnectionError:
        raise Exception("Erro de conexão ao tentar editar o título.")


def editar_autor(livro_id: int, autor: str):
    try:
        response = requests.patch(f"{API_URL}/livros/{livro_id}/autor?autor={autor}")
        return _tratar_resposta(response)
    except requests.exceptions.ConnectionError:
        raise Exception("Erro de conexão ao tentar editar o título.")


def editar_quantidade(livro_id: int, quantidade: int):
    try:
        response = requests.patch(
            f"{API_URL}/livros/{livro_id}/quantidade?quantidade={quantidade}"
        )
        return _tratar_resposta(response)
    except requests.exceptions.ConnectionError:
        raise Exception("Erro de conexão ao tentar editar o título.")


def editar_ano(livro_id: int, ano: int):
    try:
        response = requests.patch(
            f"{API_URL}/livros/{livro_id}/ano?ano={ano}", timeout=10
        )
        return _tratar_resposta(response)
    except requests.exceptions.ConnectionError:
        raise Exception("Erro de conexão ao tentar editar o ano.")


def editar_preco(livro_id: int, preco: float):
    try:
        response = requests.patch(
            f"{API_URL}/livros/{livro_id}/preco?preco={preco}", timeout=10
        )
        return _tratar_resposta(response)
    except requests.exceptions.ConnectionError:
        raise Exception("Erro de conexão ao tentar editar o preço.")


# --- Delete (Remover) ---
def deletar_livro(livro_id: int):
    try:
        response = requests.delete(f"{API_URL}/livros/{livro_id}", timeout=10)
        return _tratar_resposta(response)
    except requests.exceptions.ConnectionError:
        raise Exception("Erro de Conexão ao tentar deletar.")
