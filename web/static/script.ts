document.addEventListener("DOMContentLoaded", () => {
    const task1Form = document.querySelector<HTMLFormElement>("#task1-form");
    const task2Form = document.querySelector<HTMLFormElement>("#task2-form");

    if (task1Form) {
        task1Form.addEventListener("submit", (event) => {
            const input1 = task1Form.querySelector<HTMLInputElement>("#input1");
            const input2 = task1Form.querySelector<HTMLInputElement>("#input2");

            if (input1 && input2 && (input1.value.trim() === "" || input2.value.trim() === "")) {
                event.preventDefault();
                alert("Please fill in all fields for Task 1.");
            }
        });
    }

    if (task2Form) {
        task2Form.addEventListener("submit", (event) => {
            const input1 = task2Form.querySelector<HTMLInputElement>("#input1");

            if (input1 && input1.value.trim() === "") {
                event.preventDefault();
                alert("Please fill in the field for Task 2.");
            }
        });
    }
});
