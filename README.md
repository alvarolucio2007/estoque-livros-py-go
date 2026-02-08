# estoque-livros-py
Estoque de Livros, escrito em Python, fazendo parte de um projeto de 6 Semanas.
Este projeto é uma refatoração de um projeto da faculdade.

Para rodar este projeto (sem docker, por enquanto) voce precisará de 2 terminais:
O primeiro, crie um venv, instale o streamlit e pandas, e rode `streamlit run app.py` (o app.py é apenas para rodar o programa.)
O segundo, use e ative o venv  do primeiro, instale o uvicorn, garanta que está na pasta /src, e rode `uvicorn rotas_api:app --reload`.
No futuro será colocado um docker. Provavelmente junto a integração com Postgresql
