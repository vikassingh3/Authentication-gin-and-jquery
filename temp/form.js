// Document is ready
$(document).ready(function () {

    var x
    // Validate Username
    $('#name').keyup(function () {
        ValidateName();

    });
    function ValidateName() {
        let name = $('#name').val();
        if (name.length == "") {
            $('#namecheck').show();
            $("#namecheck").html("Name field is required")
            $("#namecheck").css("color", "red");
            $("#name").css("borderColor", "red")
            return false;
        } else if (!(/^[a-zA-Z" "]/g.test(name))) {
            $('#namecheck').show();
            $('#namecheck').html("Name contains characters only");
            $("#namecheck").css("color", "red");
            $("#name").css("borderColor", "red");
            return false;
        }
        else {
            $('#namecheck').hide();
            $("#name").css("color", "green");
            $("#name").css("borderColor", "green");
        }
        return true
    }

    // Validate Email

    $("#emails").keyup(function () {
        ValidateEmail();

    })
    function ValidateEmail() {
        const email = $("#emails").val();
        if (email == "") {
            $("femail").show();
            $("#femail").html("Email field is required")
            $("#femail").css("color", "red")
            $("#emails").css("borderColor", "red")
            return false;
        } else if (!(/^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(email))) {
            $("#femail").show();
            $("#femail").html("Please enter valid email id");
            $("#femail").css("color", "red")
            $("#emails").css("borderColor", "red")
            return false
        } else {
            $("#femail").hide();
            $("#emails").css("color", "green")
            $("#emails").css("borderColor", "green")
        }
        return true

    }



    $("#dob").change(function () {
        ValidateDOB();


    })
    function ValidateDOB() {
        var userInput = $("#dob").val();
        var dob = new Date(userInput);
        var month = Date.now() - dob.getTime();
        var Age = new Date(month);
        var year = Age.getUTCFullYear();
        var totalAge = Math.abs(year - 1970);
        var TodayDate = new Date();

        if (userInput == "") {
            $("#fdob").show();
            $("#fdob").html("Date-of-birth is required");
            $("#fdob").css("color", "red");
            $("#dob").css("borderColor", "red");
            $("#dob").css("color", "red");
            $("#age").hide()
            return false
        } else if (totalAge < 18) {
            $("#fdob").show();
            $("#fdob").html("Age should be greater than 18");
            $("#fdob").css("color", "red");
            $("#dob").css("borderColor", "red");
            $("#dob").css("color", "red");
            $("#age").hide()
            return false
        } else if (dob > TodayDate) {
            $("#fdob").show();
            $("#fdob").html("Given year can not greater than present year");
            $("#fdob").css("color", "red");
            $("#dob").css("borderColor", "red");
            $("#dob").css("color", "red");
            $("#age").hide()
            return false
        } else {
            x = totalAge
            $("#fdob").hide();
            $("#dob").css("color", "green");
            $("#dob").css("borderColor", "green");
            $("#age").show();
            $("#age").html("Age " + totalAge)
        }
        return true
    }

    $("#mobile").keyup(function () {
        ValidateMobile();

    })

    function ValidateMobile() {
        var mobile = $("#mobile").val();

        if (mobile == "") {
            $("#fmobile").show();
            $("#fmobile").html("Mobile number is required");
            $("#fmobile").css("color", "red");
            $("#mobile").css("borderColor", "red");
            $("#mobile").css("color", "red");
        } else if (mobile.length < 10) {
            $("#fmobile").show();
            $("#fmobile").html("Length of mobile number should be 10")
            $("#mobile").css("borderColor", "red");
            $("#mobile").css("color", "red");

        } else if (isNaN(mobile)) {
            $("#fmobile").show();
            $("#fmobile").html("Only digits are allowed");
            $("#fmobile").css("color", "red");
            $("#mobile").css("borderColor", "red");
            $("#mobile").css("color", "red");
        } else {
            $("#fmobile").hide();
            $("#mobile").css("borderColor", "green");
            $("#mobile").css("color", "green");

        }

        return true
    }



    // Validate Password
    $('#pass').keyup(function () {
        ValidatePassword();

    });
    function ValidatePassword() {
        let password = $('#pass').val();
        if (password == "") {
            $('#fpass').show();
            $("#fpass").html("Password field is required");
            $("#fpass").css("color", "red");
            $("#pass").css("color", "red");
            $("#pass").css("borderColor", "red")
            return false;
        } else if (!(/(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[!@#$%^&*()+=-\?;,./{}|\":<>\[\]\\\' ~_]).{8,20}/).test(password)) {
            $("#fpass").show();
            $('#fpass').html("Password length should be greater than 8 and less than 20");
            $('#fpass').css("color", "red");
            $("#pass").css("color", "red");
            $("#pass").css("borderColor", "red")
            return false;
        } else {
            $('#fpass').hide();
            $("#pass").css("borderColor", "green");
            $("#pass").css("color", "green");
        }
        return true

    }

    // Validate Confirm Password


    $('#confirm').keyup(function () {
        ValidateConfirmPasswrd();
    });
    function ValidateConfirmPasswrd() {
        let ConfPass = $('#confirm').val();
        let passwrd = $('#pass').val();
        if (passwrd != ConfPass) {
            $('#fconfirm').show();
            $('#fconfirm').html(" Password didn't Match");
            $('#fconfirm').css("color", "red");
            $("#confirm").css("color", "red");
            $("#confirm").css("borderColor", "red")
            return false;
        } else if (ConfPass == "") {
            $("#fconfirm").show();
            $("#fconfirm").html("Please confirm your password!");
            $("#fconfirm").css("color", "red");
            $("#confirm").css("color", "red");
            $("#confirm").css("borderColor", "red");
        } else {
            $('#fconfirm').hide();
            $("#confirm").css("color", "green");
            $("#confirm").css("borderColor", "green")
        }
        return true

    }
    $.Country = function () {

        var CountryObject = {
            "Australia": {
                "Queensland": ['City1', 'City2', "City3"],
                "South Australia": ['City4', 'City5', 'City6']

            },
            "India": {
                "Uttar Pradesh": ['raebareli', 'Noida', 'Lucknow'],
                "Madhya Pradesh": ['bhopal', 'jabalpur', 'Gwalior']
            }

        }

        $(window).on('load', function func() {
            var SelCountry = ($('#country').get(0));
            var SelState = ($('#state').get(0));
            for (var country in CountryObject) {
                SelCountry.options[SelCountry.options.length] = new Option(country);
            }
            $("#country").change(function () {
                SelState.length = 1;
                if (this.selectedIndex < 1) return
                for (var state in CountryObject[this.value]) {
                    SelState.options[SelState.options.length] = new Option(state);
                }
            })
            SelCountry.change()
            // return true
        })

    }
    $.Country();

    $("#country").change(function () {
        ValidateCountry();
    })

    function ValidateCountry() {
        var count = $("#country").val();
        $("#state").val();

        if (count == 0) {
            $("#fcountry").show();
            $("#state").css("color", "");
            $("#state").css("borderColor", "")
            $("#fstate").html("");
            $("#fcountry").html("Country is required");
            $("#fcountry").css("color", "red")
            $("#country").css("borderColor", "red")
            $("#country").css("color", "red")

        }
        else {
            $("#fcountry").hide();
            $("#state").css("color", "");
            $("#state").css("borderColor", "")
            $("#fstate").html("");
            $("#country").css("borderColor", "green")
            $("#country").css("color", "green")

        }
        return true
    }


    $("#state").change(function () {
        ValidateState();
    })

    function ValidateState() {
        var state = $("#state").val();

        if (state == 0) {
            $("#fstate").show();
            $("#fstate").html("State is required");
            $("#fstate").css("color", "red")
            $("#state").css("borderColor", "red")
            $("#state").css("color", "red")
            return false
        }
        else {
            $("#fstate").hide();
            $("#state").css("borderColor", "green")
            $("#state").css("color", "green")

        }
        return true
    }

    $(function () {
        var dtToday = new Date();

        var month = dtToday.getMonth() + 1;
        var day = dtToday.getDate();
        var year = dtToday.getFullYear();

        if (month < 10)
            month = '0' + month.toString();
        if (day < 10)
            day = '0' + day.toString();

        var maxDate = year + '-' + month + '-' + day;
        $('#dob').attr('max', maxDate);
    });



    // Submit button
    $('#submitbtn').click(function () {
        ValidateName();
        ValidateEmail();
        ValidateDOB();
        ValidatePassword();
        ValidateConfirmPasswrd();
        if (
            ValidateName() == true &&
            ValidateEmail() == true &&
            ValidatePassword() == true &&
            ValidateDOB() == true &&
            ValidateConfirmPasswrd() == true

        ) {

            var name = $("#name").val();
            var email = $("#emails").val();
            var gender = $("#gender").val();
            var address = $("#address").val();
            var dob = $("#dob").val();
            var country = $("#country").val();
            var state = $("#state").val();
            var password = $("#pass").val();
            var tok3 = localStorage.getItem("Bearer Token")

            url = "http://127.0.0.1:10000/admin/newdata";
            params = {
                method: "post",
                headers: {
                    'Content-Type': 'application/json',
                    Authorization: `Bearer ${tok3}`
                },
                body: JSON.stringify({
                    name,
                    email,
                    gender,
                    address,
                    dob,
                    country,
                    state,
                    password
                }),
            }
            fetch(url, params).then(res => {
                res.json()
            })
                .then(data => console.log(data))
            window.location.href = "http://127.0.0.1:10000/data/form3.html"

            $("#state option").each(function () {
                $(this).remove();
            });
            $('#state').append($('<option>Select State</option>'));

            $("#age").html("");
            $("#fsubmit").html("");


            // $("select#state option[value='state']").remove();

        } else {
            $("#display_message").html("");
            $("#display_message1").html("");
            $("#display_message2").html("");
            $("#display_message3").html("");
            $("#display_message6").html("");
            $("#display_message4").html("");
            $("#display_message5").html("");
            $("#state").css("color", "red");
            $("#state").css("borderColor", "red");
            $("#fsubmit").html("All fields are required");
            $("#fsubmit").css("color", "red");
        }
    });
});
