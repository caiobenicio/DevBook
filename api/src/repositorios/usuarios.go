package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositorio de usuarios
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuario no bd
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	//statement, erro := repositorio.db.Prepare(
	//	"INSERT INTO public.usuarios(nome, nick, email, senha) VALUES ($1, $2, $3, $4)")
	var id int
	if erro := repositorio.db.
		QueryRow("INSERT INTO public.usuarios(nome, nick, email, senha, criadoem) VALUES ($1, $2, $3, $4, now()) RETURNING ID",
			usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha).Scan(&id); erro != nil {
		return 0, erro
	}

	return uint64(id), nil
}

// Buscar traz todos os usuarios que atendem um filtro de nome ou nick
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) //%nomeOuNick%

	linhas, erro := repositorio.db.
		Query("select id, nome, nick, email, criadoem from usuarios where nome LIKE $1 or nick LIKE $2",
			nomeOuNick, nomeOuNick)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarPorID traz todos os usuarios que atendem um filtro de nome ou nick
func (repositorio Usuarios) BuscarPorID(ID uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query("select id, nome, nick, email, criadoem from usuarios where id = $1", ID)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario
	if linhas.Next() {
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return modelos.Usuario{}, erro
		}
	}
	return usuario, nil
}

// Atualizar editar um usuario
func (repositorio Usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
	statement, erro := repositorio.db.Prepare("update usuarios set nome = $1, nick = $2, email = $3 where id = $4")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	}

	return nil
}

// Deletar apaga um usuario do bd
func (repositorio Usuarios) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from usuarios where id = $1")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linha, erro := repositorio.db.Query("select id, senha from usuarios where email = $1", email)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

// Seguir permite que um usuario siga outro
func (repositorio Usuarios) Seguir(usuarioID, seguidorID uint64) error {
	linha, erro := repositorio.db.Query("insert into seguidores(usuario_id, seguidor_id) values ($1, $2) ON CONFLICT (usuario_id, seguidor_id) DO NOTHING",
		usuarioID, seguidorID)
	if erro != nil {
		return erro
	}
	defer linha.Close()

	return nil
}

// PararDeSeguir permite que um usuario pare de seguir outro
func (repositorio Usuarios) PararDeSeguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from seguidores where usuario_id=$1 and seguidor_id=$2")
	if erro != nil {
		return erro
	}

	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

// BuscarSeguidores retorna os seguidores de um usuario
func (repositorio Usuarios) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
					select u.id, u.nome, u.nick, u.email, u.criadoem
					from usuarios u inner join seguidores s on u.id = s.seguidor_id 
					where s.usuario_id = $1`, usuarioID)
	if erro != nil {
		return []modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var seguidores []modelos.Usuario
	for linhas.Next() {
		var seguidor modelos.Usuario

		if erro := linhas.Scan(&seguidor.ID, &seguidor.Nome, &seguidor.Nick, &seguidor.Email, &seguidor.CriadoEm); erro != nil {
			return nil, erro
		}

		seguidores = append(seguidores, seguidor)
	}

	return seguidores, nil
}

// BuscarSeguindo retorna os seguidores de um usuario
func (repositorio Usuarios) BuscarSeguindo(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
					select u.id, u.nome, u.nick, u.email, u.criadoem
					from usuarios u inner join seguidores s on u.id = s.usuario_id 
					where s.seguidor_id = $1`, usuarioID)
	if erro != nil {
		return []modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var seguidores []modelos.Usuario
	for linhas.Next() {
		var seguidor modelos.Usuario

		if erro := linhas.Scan(&seguidor.ID, &seguidor.Nome, &seguidor.Nick, &seguidor.Email, &seguidor.CriadoEm); erro != nil {
			return nil, erro
		}

		seguidores = append(seguidores, seguidor)
	}

	return seguidores, nil
}

// BuscarSenha traz a senha de um usuario pelo ID
func (repositorio Usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linha, erro := repositorio.db.Query("select senha from usuarios where id = $1", usuarioID)
	if erro != nil {
		return "", nil
	}
	defer linha.Close()

	var usuario modelos.Usuario
	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", nil
		}
	}

	return usuario.Senha, nil
}

// AtualizarSenha altera a senha de um usuario
func (repositorio Usuarios) AtualizarSenha(usuarioID uint64, senha string) error {
	statement, erro := repositorio.db.Prepare("update usuarios set senha = $1 where id = $2")
	if erro != nil {
		return nil
	}
	defer statement.Close()

	if _, erro = statement.Exec(senha, usuarioID); erro != nil {
		return erro
	}

	return nil
}
