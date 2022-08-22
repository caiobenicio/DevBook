package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

// NovoRepositorioDepublicacaos cria um repositorio de publicações
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repositorio Publicacoes) Criar(pub modelos.Publicacao) (uint64, error) {
	var id int
	if erro := repositorio.db.
		QueryRow("INSERT INTO publicacoes(titulo, conteudo, autor_id) VALUES ($1, $2, $3) RETURNING ID",
			pub.Titulo, pub.Conteudo, pub.AutorID).Scan(&id); erro != nil {
		return 0, erro
	}

	return uint64(id), nil
}

// BuscarPorID traz uma publica que atendem ao ID
func (repositorio Publicacoes) BuscarPorID(ID uint64) (modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
					select p.*, u.nick from publicacoes p inner join usuarios u
					on u.id = p.autor_id where p.id =$1`, ID)
	if erro != nil {
		return modelos.Publicacao{}, erro
	}
	defer linhas.Close()

	var publicacao modelos.Publicacao
	if linhas.Next() {
		if erro := linhas.Scan(&publicacao.ID, &publicacao.Titulo, &publicacao.Conteudo, &publicacao.AutorID,
			&publicacao.Curtidas, &publicacao.CriadaEm, &publicacao.AutorNick); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}

	if publicacao.AutorID != 0 {

		return publicacao, nil
	}

	return publicacao, nil
}

// Buscar traz as publicações dos usuários e do proprio usuário
func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.
		Query(`select distinct p.*, u.nick from publicacoes p
				inner join usuarios u on u.id = p.autor_id
				inner join seguidores s on p.autor_id = s.usuario_id
				where u.id = $1 or s.seguidor_id =$2 order by 1 desc`,
			usuarioID, usuarioID)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao
	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro := linhas.Scan(&publicacao.ID, &publicacao.Titulo, &publicacao.Conteudo,
			&publicacao.AutorID, &publicacao.Curtidas, &publicacao.CriadaEm,
			&publicacao.AutorNick); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// Atualizar Altera os dados de uma publicação
func (repositorio Publicacoes) Atualizar(ID uint64, publicacao modelos.Publicacao) error {
	statement, erro := repositorio.db.Prepare("update publicacoes set titulo = $1, conteudo = $2 where id = $3")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, ID); erro != nil {
		return erro
	}

	return nil
}

// Deletar apaga um publicação do bd
func (repositorio Publicacoes) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from publicacoes where id = $1")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// Buscar traz as todas as publicações de um usuário
func (repositorio Publicacoes) BuscarPorUsuario(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.
		Query(`select p.*, u.nick from publicacoes p
				inner join usuarios u on u.id = p.autor_id
				where p.autor_id = $1`,
			usuarioID)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao
	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro := linhas.Scan(&publicacao.ID, &publicacao.Titulo, &publicacao.Conteudo,
			&publicacao.AutorID, &publicacao.Curtidas, &publicacao.CriadaEm,
			&publicacao.AutorNick); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// Curtir adiciona uma curtida na publicação
func (repositorio Publicacoes) Curtir(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare("update publicacoes set curtidas = curtidas + 1 where id = $1")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}

// Descurtir subtrai uma curtida na publicação
func (repositorio Publicacoes) Descurtir(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare(`
		update publicacoes set curtidas = 
		CASE WHEN curtidas > 0  THEN curtidas - 1 
			 ELSE 0 END 
		where id = $1`)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}
