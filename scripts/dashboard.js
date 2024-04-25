function toggleStatus(cell) {
   if (cell.classList.contains("yes")) {
      cell.classList.remove("yes");
      cell.classList.add("maybe");
      cell.textContent = "Maybe";
   } else if (cell.classList.contains("maybe")) {
      cell.classList.remove("maybe");
      cell.classList.add("no");
      cell.textContent = "NO";
   } else if (cell.classList.contains("no")) {
      cell.classList.remove("no");
      cell.classList.add("yes");
      cell.textContent = "YES";
   } else {
      cell.classList.add("yes");
      cell.textContent = "YES";
   }
}

function toggleStatusCol(col_idx) {
    for (let i = 0; i < 5; i ++) {
        toggleStatus(document.getElementById(`${i}-${col_idx}`));
    }
}
