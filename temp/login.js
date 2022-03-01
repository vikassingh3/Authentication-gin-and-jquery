function ValidateName() {
    var username = document.getElementById("username").value;
    if (username == "") {
        document.getElementById("fusername").innerHTML = "plaese provide username";
        document.getElementById("fusername").style.color = "red"
        return false
    } else {
        document.getElementById("fusername").innerHTML = "";
        document.getElementById("fusername").style.color = "red"
    }
    return true
}
function ValidatePassword() {
    var username = document.getElementById("password").value;
    if (username == "") {
        document.getElementById("fpassword").innerHTML = "plaese provide password";
        document.getElementById("fpassword").style.color = "red"
        return false
    } else {
        document.getElementById("fpassword").innerHTML = "";
        document.getElementById("fpassword").style.color = "red"
    }
    return true
}

function SignIn() {
    ValidateName();
    ValidatePassword();
    if (ValidateName() == true && ValidatePassword() == true) {
        var Email = document.getElementById("username").value;
        var password = document.getElementById("password").value;


        fetch("http://127.0.0.1:10000/login", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                Email,
                password
            }),
        })
            .then((response) => {
                if (response.status !== 200) {
                    console.log(response)
                    console.log("error")
                } else {
                    response.json()
                        .then(data => {
                            console.log('Success:', data);
                            localStorage.setItem("Bearer Token", data.access_token)

                        })
                    window.location.href = "http://127.0.0.1:10000/data/form3.html"
                }
            })


    }
}
function ResetPass() {
    fetch("http://127.0.0.1:10000/link", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        }
    }).then(res => {
        return res.json()
    }).then(data => console.log(data) )
}

function forgot(){
    window.location.href = "http://127.0.0.1:10000/data/email.html"
}