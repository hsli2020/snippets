// (A) INITIALIZE - DOUBLE CLICK TO EDIT CELL
window.addEventListener("DOMContentLoaded", () => {
  for (let cell of document.querySelectorAll(".editable td")) {
    cell.ondblclick = () => editable.edit(cell);
  }
});

var editable = {
  // (B) "CONVERT" TO EDITABLE CELL
  edit : cell => {
    // (B1) REMOVE "DOUBLE CLICK TO EDIT"
    cell.ondblclick = "";

    // (B2) EDITABLE CONTENT
    cell.contentEditable = true;
    cell.focus();

    // (B3) "MARK" CURRENT SELECTED CELL
    cell.classList.add("edit");
    editable.selected = cell;

    // (B4) PRESS ENTER OR CLICK OUTSIDE TO END EDIT
    window.addEventListener("click", editable.close);
    cell.onkeydown = evt => { if (evt.key=="Enter") {
      editable.close(true);
      return false;
    }};
  },

  // (C) END "EDIT MODE"
  selected : null,  // current selected cell
  close : evt => { if (evt.target != editable.selected) {
    // (C1) DO WHATEVER YOU NEED
    // check value?
    // send value to server?
    // update calculations in table?

    // (C2) REMOVE "EDITABLE"
    window.getSelection().removeAllRanges();
    editable.selected.contentEditable = false;

    // (C3) RESTORE CLICK LISTENERS
    window.removeEventListener("click", editable.close);
    let cell = editable.selected;
    cell.ondblclick = () => editable.edit(cell);

    // (C4) "UNMARK" CURRENT SELECTED CELL
    editable.selected.classList.remove("edit");
    editable.selected = null;
  }}
};