document.addEventListener("DOMContentLoaded", () => {
    const generateForm = document.getElementById("generate-form");
    const asciiArtInput = document.getElementById("ascii-art-input");
    const generateBtn = document.getElementById("generate-btn");
    const clearBtn = document.getElementById("clear-btn");
    const asciiArtContainer = document.getElementById("ascii-art-container");

    generateForm.addEventListener("submit", (e) => {
        e.preventDefault();

        const formData = new FormData(generateForm);
        const word = formData.get("word");
        const font = formData.get("font");

        generateAsciiArt(word, font);
    });

    clearBtn.addEventListener("click", () => {
        asciiArtInput.value = "";
        asciiArtContainer.textContent = "";
    });

    function generateAsciiArt(word, font) {
        fetch("/generate", {
            method: "POST",
            body: JSON.stringify({ word, font }),
            headers: {
                "Content-Type": "application/json",
            },
        })
            .then((response) => response.text())
            .then((result) => {
                asciiArtContainer.textContent = result;
            })
            .catch((error) => {
                console.error("Failed to generate ASCII art:", error);
            });
    }
});
