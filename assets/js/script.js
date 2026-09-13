// Delete the bracket
const deleteButton = document.getElementById("delete-bracket");

deleteButton.addEventListener("click", () => {
    const bracket = document.getElementById("bracket");

    if (bracket) {
        bracket.remove();
    }
});