import pandas as pd
import streamlit as st

import src.servico as srv


@st.cache_resource
def get_service():
    return srv.Service()


class FrontEnd:
    def __init__(self) -> None:
        self.servico = get_service()

    def exibir_mensagens_pendentes(self) -> None:
        if "success_message" in st.session_state:
            _ = st.success(st.session_state["success_message"])
            del st.session_state["success_message"]

    def renderizar_menu_lateral(self) -> str:
        with st.sidebar:
            _ = st.header("Selecione a ação")
            opcao = st.radio(
                "",
                (
                    "Cadastrar Livro",
                    "Listar Livros",
                    "Buscar Livros",
                    "Atualizar Livros",
                    "Excluir Livros",
                    "Gerar Relatórios",
                ),
            )
            return opcao

    def renderizar_cadastro(self) -> None:
        self.exibir_mensagens_pendentes()
        _ = st.markdown("Cadastro de novo livro")
        with st.form("form_cadastro"):
            titulo = st.text_input("Título do Livro", max_chars=100)
            col1, col2 = st.columns(2)
            with col1:
                autor = st.text_input("Autor do Livro", max_chars=100)
            with col2:
                ano = st.number_input(
                    "Ano de lançamento", min_value=0, format="%d", step=1
                )
            col3, col4 = st.columns(2)
            with col3:
                preco = st.number_input(
                    "Preço do Livro", min_value=0.01, format="%.2f", step=0.01
                )
            with col4:
                quantidade = st.number_input(
                    "Quantidade do Livro", min_value=0, format="%d", step=1
                )
            clique_salvar = st.form_submit_button("Cadastrar e Salvar", width="stretch")
            if clique_salvar:
                try:
                    self.servico.cadastrar_livro(
                        titulo=str(titulo),
                        preco=float(preco),
                        autor=str(autor),
                        ano=int(ano),
                        quantidade=int(quantidade),
                    )
                    st.session_state["success_message"] = (
                        f"Produto {titulo} adicionado com sucesso!"
                    )
                    st.rerun()
                except ValueError as e:
                    _ = st.error(f"Falha no cadastro! {e}")
