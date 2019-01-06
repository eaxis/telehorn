$(function(){
    const SUCCESS = 'success';
    const WARNING = 'warning';

    $('.custom-alert .close').on('click', hideAlert);
    $('button.send').on('click', send);
    $('button.submit').on('click', submit);

    function send(e)
    {
        if (
            !$('.token').val().length ||
            !$('.chats').val().length ||
            !$('.message').val().length
        )
        {
            showAlert(WARNING, 'You should fill in all fields.');
            return;
        }

        $('.preview-payload').html(
            $('.message').val()
        );

        $('.modal').modal('show');
    }

    function submit(e)
    {
        var request = {
            token: "",
            chats: [],
            message: ""
        };

        request.token = $('.token').val();
        request.chats = filterChats($('.chats').val());
        request.message = $('.message').val();

        var promise = $.ajax({
            url: '/submit',
            type: 'POST',
            dataType: "json",
            contentType: "application/json; charset=utf-8",
            data: JSON.stringify(request)
        });

        promise.always(function (data) {
            $('.modal').modal('hide');

            if (data.success) {
                disableInterface(true);

                setTimeout(function () {
                    disableInterface(false);
                }, 10000);

                showAlert(SUCCESS, data.description);
                return;
            }

            if (data.responseJSON && !data.responseJSON.success) {
                showAlert(WARNING, data.responseJSON.description);
                return;
            }

            showAlert(WARNING, 'Unsuccessful request.');
        });
    }

    function hideAlert(e)
    {
        alert = $('.custom-alert');

        alert.fadeOut({
            done: function () {
                alert.removeClass('alert-warning');
                alert.removeClass('alert-success');
            }
        });
    }

    function showAlert(type, text)
    {
        alert = $('.custom-alert');

        alert.find('.alert-text').text(text);
        alert.find('.alert-title').text(type === SUCCESS ? 'Success!' : 'Warning!');

        alert.addClass('alert-' + type).fadeIn();
    }

    function disableInterface(disable)
    {
        let elements = [
            'token', 'chats',
            'message', 'send',
            'submit'
        ];

        for (index in elements) {
            $('.' + elements[index]).attr('disabled', disable);
        }
    }

    function filterChats(chats)
    {
        chats = chats.split(',');

        chats = $.map(chats, function (element) {
           return parseInt(element)
        });

        return chats;
    }
});