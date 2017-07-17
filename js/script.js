function w3_open() {
  document.getElementsByClassName("w3-sidenav")[0].style.display = "block";
}

function w3_close() {
  document.getElementsByClassName("w3-sidenav")[0].style.display = "none";
}

/*test json file of database json return*/

var jsonreturn = `\
[\
    {\
        "id":1,\
        "username":"test",\
        "email":"test@mail.com",\
    },\
    {\
        "id":2,\
        "username":"test",\
        "email":"test@mail.com",\
    },\
    {\
        "id":3,\
        "username":"test",\
        "email":"test@mail.com",\
    },\
]`
