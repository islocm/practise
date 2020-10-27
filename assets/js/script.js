function searchp() {
    var txtValue, tr, input, td, filter;
    input = document.getElementById("searchp");
    table = document.getElementById("mytable");
    tr = table.getElementsByTagName("tr");
    filter = input.value.toUpperCase();
    for (let i = 0; i < tr.length; i++) {
        td = tr[i].getElementsByTagName("td")[0];
        if (td) {
            txtValue = td.textContent || td.innerText;
            if (txtValue.toUpperCase().indexOf(filter) > -1) {
                tr[i].style.display = "block";
            } else {
                tr[i].style.display = "none";
            }
        }

    }

}

function searchc() {
    var txtValue, tr, input, td, filter;
    input = document.getElementById("searchc");
    table = document.getElementById("mytable");
    tr = table.getElementsByTagName("tr");
    filter = input.value.toUpperCase();
    for (let i = 0; i < tr.length; i++) {
        td = tr[i].getElementsByTagName("td")[1];
        if (td) {
            txtValue = td.textContent || td.innerText;
            if (txtValue.toUpperCase().indexOf(filter) > -1) {
                tr[i].style.display = "block";
            } else {
                tr[i].style.display = "none";
            }
        }

    }

}

function searchm() {
    var txtValue, tr, input, td, filter;
    input = document.getElementById("searchm");
    table = document.getElementById("mytable");
    tr = table.getElementsByTagName("tr");
    filter = input.value.toUpperCase();
    for (let i = 0; i < tr.length; i++) {
        td = tr[i].getElementsByTagName("td")[2];
        if (td) {
            txtValue = td.textContent || td.innerText;
            if (txtValue.toUpperCase().indexOf(filter) > -1) {
                tr[i].style.display = "block";
            } else {
                tr[i].style.display = "none";
            }
        }

    }

}

function searcho() {
    var txtValue, tr, input, td, filter;
    input = document.getElementById("searcho");
    table = document.getElementById("mytable");
    tr = table.getElementsByTagName("tr");
    filter = input.value.toUpperCase();
    for (let i = 0; i < tr.length; i++) {
        td = tr[i].getElementsByTagName("td")[3];
        if (td) {
            txtValue = td.textContent || td.innerText;
            if (txtValue.toUpperCase().indexOf(filter) > -1) {
                tr[i].style.display = "block";
            } else {
                tr[i].style.display = "none";
            }
        }

    }

}