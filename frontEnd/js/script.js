function w3_open() {
  document.getElementsByClassName("w3-sidenav")[0].style.display = "block";
}

function w3_close() {
  document.getElementsByClassName("w3-sidenav")[0].style.display = "none";
}

function openPage(name) {
  var i;
  var x = document.getElementsByClassName("page");
  for (i = 0; i < x.length; i++) {
    x[i].style.display = "none";
  }
  document.getElementById(name).style.display = "block";
}