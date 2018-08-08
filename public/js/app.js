$(document).ready(function () {

    // The connection function
    $("#connection").click(function (e) {
        e.preventDefault();
        // var button = $(this);
        // button.attr("disabled", true);

        $.ajax({
            type: "GET",
            url: '/.netlify/functions/connection',
            contentType: 'application/json',
            crossDomain: true,
            success: function (data) {
                swal("Success!", "Response: " + data["message"], "success"); //This prints the response with the header.

            },
            error: function () {
                alert('fail');

            }
        });

    })

    // Form ajax function
    $("#ajax").submit(function (e) {

        e.preventDefault();

        // var button = $(this).find('#submit');
        // button.attr("disabled", true);

        var $form = $(this);
        $.post(
            $form.attr("action"),
            $form.serialize()
        ).done(function (data) {
            console.log(data)
            alert("Hello, " + data["name"][0])
        });

    })

    // GET ENVIROMENT
    $("#envButton").click(function (e) {
        e.preventDefault();
        // var button = $(this);
        // button.attr("disabled", true);

        $.ajax({
            type: "GET",
            url: '/.netlify/functions/env',
            contentType: 'application/json',
            crossDomain: true,
            success: function (data) {
                console.log(data)
                console.log(data["ENV_KEY"])
                var env = data["ENV_KEY"]
                $("#envButton").data('env', env)
                var modal = $('#envModal')
                modal.find('.modal-body').text(env)
                modal.modal('toggle')
            },
            error: function () {
                alert('fail');

            }
        });

    })

    // GET Cyrpto pricing
    $("#cyrptoForm").submit(function (e) {
        
        e.preventDefault();

        // var button = $(this);
        // button.attr("disabled", true);

        var $form = $(this);
        $.post(
            $form.attr("action"),
            $form.serialize()
        ).done(function (data) {
            console.log(data)
            console.log(data["Name"])
            var name = data["Name"]
            var symbol = data["Symbol"]
            var rank = data["Rank"]
            var price = data["Quote"]["Price"]

            var modal = $('#cryptoModal')
            modal.find('.modal-title').text(name)
            modal.find('.modal-body').html("<h3> $" + price + "</h3>")
            modal.modal('toggle')
        });

    })
});