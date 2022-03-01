$(document).ready(function () {

    var tok1 = localStorage.getItem("Bearer Token");

    fetch("http://127.0.0.1:10000/admin/all", {
        method: 'GET',
        headers: { Authorization: `Bearer ${tok1}` },
    })
        .then((response) => {
            return response.json();
            // console.log(response)
        })
        .then((data) => {
            console.log(data)
            var admin = data
            let Alldaata = `   
                
                    <tr>
                    <th>ID</th>
                    <th>Name</th>
                    <th>Email</th> 
                    <th>Address</th> 
                    <th>Gender</th>
                    <th>Country</th>
                    <th>Action</th>
                    <th>Action</th>
                    </tr>  
            `;
            for (let x of admin.data) {
                Alldaata +=
                    `<tr>
                <td class="nr">${x.ID}</td>
                <td>${x.name} </td>
                <td>${x.email}</td>
                <td>${x.address}</td>
                <td>${x.gender}</td>
                <td>${x.country}</td>

                <td> <br> <button class="updatedata" type="button" > edit </button> <br><br>  </td>
                <td> <br><button class="use-address" type="button" > Delete </button> </td>`
            }

            document.getElementById("choose-address-table").innerHTML = Alldaata;


            $(".use-address").click(function () {
                var txt;
                var id = $(this).closest("tr").find(".nr").text();
                if (confirm("Are you sure!") == true) {
                    var xid = id
                    fetch(`http://127.0.0.1:10000/admin/delete/${id}`, {
                        method: 'DELETE',
                    })
                    window.location.href = "http://127.0.0.1:10000/data/form3.html";
                } else {
                    txt = "press cancel"
                }
            })


            $(".updatedata").click(function () {
                document.getElementById("choose-address-table").innerHTML = "";
                document.getElementById("btn").style.display = "none";
                document.getElementById("logout").style.display = "none";




                document.getElementById("datails").innerHTML = ` <form name="myform" id="myform"> 
                <div class="form-group" >
                     <label >Name</label>
                     <input type="text" id="name" name="name" class="form-control" placeholder="name" >
                     <span id="error_name" class="text-danger" ></span>							
                </div>
                <div class="form-group">
                    <label>Email</label>
                    <input type="myEmail" id="emails" name="email"  class="form-control" placeholder="email"  >
                    <span id="emailmsg" class="text-danger"></span>							
                </div>
                <div class="form-group">
                    <label >Address</label>
                    <input type="text" id="address" name="address" class="form-control" placeholder="address" >
                    <span id="address_message" class="text-danger"  ></span>							
                </div>

                <button id="datam" class="updatem" > Update </button> <br><br>`
                var id = $(this).closest("tr").find(".nr").text();

                fetch(`http://127.0.0.1:10000/admin/get/${id}`)
                    .then((response) => {
                        return response.json();
                    })
                    .then((data) => {
                        let authors = data;
                        $("#name").val(authors.data.name),
                            $("#emails").val(authors.data.email),
                            $("#address").val(authors.data.address),
                            // console.log(authors)

                            $(".updatem").click(function () {
                                fetch(`http://127.0.0.1:10000/admin/update/${id}`, {
                                    method: 'PUT',
                                    headers: {
                                        'Content-Type': 'application/json'
                                    },
                                    body: JSON.stringify({
                                        name: $("#name").val(),
                                        email: $("#emails").val(),
                                        address: $("#address").val(),
                                    }),
                                }).then(response => response.json()).then(data => console.log(data))
                            })
                    })
            })
        })
    $("#logout").click(function () {
        var tok = localStorage.getItem("Bearer Token");
        fetch("http://127.0.0.1:10000/logout", {
            type: 'POST',
            contentType: 'application/json',
            headers: { Authorization: `Bearer ${tok}` }
        }).then(res => {
            if (res.status != 200){
                console.log("error")
            }else{
                window.location.href = "http://127.0.0.1:10000/data/index.html"
            }
        })

    })


})

