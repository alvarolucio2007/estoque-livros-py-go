package test

//func TestAPI(t *testing.T) {
//	t.Run("apiListarLivros", func(t *testing.T) {
//		testarHandlerCarregarLivros(t)
//	})
//}

//func testarHandlerCarregarLivros(t *testing.T) {
//	gin.SetMode(gin.TestMode)
//	r := gin.Default()
//	r.GET("/livros", HandlerListarLivros)
//	criarLivroTeste(t)
//	url := "/livros"
//	req, _ := http.NewRequest("GET", url, nil)
//	w := httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//	if w.Code != http.StatusOK {
//		t.Errorf("esperava status 200,recebi %d", w.Code)
//	}
//}

//func testarHandlerListarID(t *testing.T) {
//	gin.SetMode(gin.TestMode)
//	r := gin.Default()
//	r.GET("/livros/listar_id", testarHandlerListarID)
//}
