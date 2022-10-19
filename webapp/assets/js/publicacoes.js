$('#nova-publicacao').on('submit', criarPublicacao);
$(document).on('click', '.curtir-publicacao', curtirPublicacao);
$(document).on('click', '.descurtir-publicacao', descurtirPublicacao);
$('#atualizar-publicacao').on('click', atualizarPublicacao);

function criarPublicacao(evento) {
    evento.preventDefault();
    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        }
    }).done(function () {
        window.location = "/home";
    }).fail(function (erro) {
        alert("Erro ao criar Publicação!");
    });
}

function curtirPublicacao(evento) {
    evento.preventDefault();
    var iconHeart = $(evento.target);
    const publicacaoId = iconHeart.closest('div').data('publicacao-id');
    iconHeart.prop('disabled', true);

    $.ajax({
        url: `/publicacoes/${publicacaoId}/curtir`,
        method: "POST"
    }).done(function () {
        $('#curtidas').text(parseInt($('#curtidas').text()) +1);
        iconHeart.addClass('descurtir-publicacao');
        iconHeart.addClass('text-danger');
        iconHeart.removeClass('curtir-publicacao');        
    }).fail(function (erro) {
        alert("Erro ao criar Publicação!");
    }).always(function() {
        iconHeart.prop('disabled', false);
    });
}

function descurtirPublicacao(evento) {
    evento.preventDefault();
    var iconHeart = $(evento.target);
    const publicacaoId = iconHeart.closest('div').data('publicacao-id');
    
    iconHeart.prop('disabled', true);
    $.ajax({
        url: `/publicacoes/${publicacaoId}/descurtir`,
        method: "POST"
    }).done(function () {
        $('#curtidas').text(parseInt($('#curtidas').text()) -1);
        iconHeart.addClass('curtir-publicacao');
        iconHeart.removeClass('text-danger');
        iconHeart.removeClass('descurtir-publicacao');        
    }).fail(function (erro) {
        alert("Erro ao criar Publicação!");
    }).always(function() {
        iconHeart.prop('disabled', false);
    });
}

function atualizarPublicacao() {
    $(this).prop('disabled', true);
    const publicacaoId = $(this).data('publicacao-id');

    $.ajax({
        url: `/publicacoes/${publicacaoId}`,
        method: "PUT",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        }
    }).done(function () {
        alert("Publicação editada com sucesso!");
    }).fail(function (erro) {
        alert("Erro ao aditar a Publicação!");
    }).always(function() {
        $('#atualizar-publicacao').prop('disabled', false);
    }); 
}