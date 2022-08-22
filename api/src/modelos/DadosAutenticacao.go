package modelos

// DadosAutenticacao contém o token eo id do usuário autenticado
type DadosAutenticacao struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
