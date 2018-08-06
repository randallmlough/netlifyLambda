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

});