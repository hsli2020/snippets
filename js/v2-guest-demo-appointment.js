jQuery(document).ready(function () {

    $("#contactUsBtn").on("click", function(e) {
        e.preventDefault();
        var input = document.getElementById("emailorphone")
        input.setCustomValidity("");
        var reg = /^[a-zA-Z0-9.!#$%&’*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/gi;
        if (!input.value.match(/[a-z]/i)) {
            reg = /^\(?(\d{3})\)?[- ]?(\d{3})[- ]?(\d{4})$/;
        }
        console.log(input.value);
        if (reg.test(input.value)==false) {
            input.setCustomValidity("welcome.demo.request.phone.or.email.wrong".format());
            input.reportValidity();
            return;
        }
        console.log("submit");
        // send email
        $("body").css("cursor", "progress");
        var $btn = $(this);
        $btn.addClass("disabled");

        var formData = new FormData(document.forms['form_appointment']);
        $.ajax({
            url : "/demoAppointment",
            data : formData,
            type : "POST",
            processData : false,
            contentType : false,
            dataType : "json",
        }).done(function(json) {
            if (json.status !== 0) {
                swal({
                    title : "welcome.demo.request.response.title.error".format(),
                    text : json.message,
                    type : "error",
                });
            } else {
                swal({
                    title :"welcome.demo.request.response.title.success".format() ,
                    text : "welcome.demo.request.response.success".format(),
                    type : "success",
                }).then(function() {
                    $('#form_demo_appointment').trigger("reset");
                });
            }
        }).fail(function(xhr, status, errorThrown) {
            swal({
                title : "welcome.demo.request.response.title.error".format(),
                text : errorThrown,
                type : "error",
            });

        }).always(function(xhr, status) {
            $("body").css("cursor", "default");
            $btn.removeClass("disabled");
        });

        return false;
    });
});
