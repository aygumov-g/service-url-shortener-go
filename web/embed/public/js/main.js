new class {
    constructor(data) {
        this.data = data;
        this.data["input"].addEventListener("input", () => {
            if (this.data["input"].value.length > 0) {
                this.showBtnContainer();
                this.showBtnText();
            } else {
                this.hideBtnContainer();
                this.hideBtnText();
            }
        });
        this.data["btnContainer"].addEventListener("click", () => {
            if (this.data["btnContainer"].classList.contains("shortener__post-button-container-show")) {
                if (this.data["btnText"].classList.contains("shortener__post-button-text-show")) {
                    if (!this.data["input"].value.trim().length === 0 || !this.data["input"].value.includes(".")) {
                        this.showErrText("Некорректная ссылка");
                    } else {
                        this.post(this.data["input"].value);
                    }
                };
            };
        });
        this.hideBtnContainer();
        this.hideBtnText();
        this.hideResultContainer();
        this.hideErrText();
    };
    showBtnContainer() {
        this.data["btnContainer"].classList.add("shortener__post-button-container-show");
        this.data["btnContainer"].classList.remove("shortener__post-button-container-hide");
    };
    hideBtnContainer() {
        this.data["btnContainer"].classList.add("shortener__post-button-container-hide");
        this.data["btnContainer"].classList.remove("shortener__post-button-container-show");
    };
    showBtnText() {
        this.data["btnText"].classList.add("shortener__post-button-text-show");
        this.data["btnText"].classList.remove("shortener__post-button-text-hide");
    };
    hideBtnText() {
        this.data["btnText"].classList.add("shortener__post-button-text-hide");
        this.data["btnText"].classList.remove("shortener__post-button-text-show");
    };
    showErrText(text) {
        this.data["errText"].classList.add("shortener__error-text-show");
        this.data["errText"].classList.remove("shortener__error-text-hide");
        this.data["errText"].textContent = text;
        this.hideResultContainer();
    };
    hideErrText() {
        this.data["errText"].classList.add("shortener__error-text-hide");
        this.data["errText"].classList.remove("shortener__error-text-show");
        this.data["errText"].textContent = "";
    };
    showResultContainer() {
        this.data["resultContainer"].classList.add("shortener__result-container-show");
        this.data["resultContainer"].classList.remove("shortener__result-container-hide");
    }
    hideResultContainer() {
        this.data["resultContainer"].classList.add("shortener__result-container-hide");
        this.data["resultContainer"].classList.remove("shortener__result-container-show");
    }
    setLink(url) {
        //this.data["linkURL"].href = url;
        this.data["linkText"].textContent = url;
    }
    async post(url) {
        try {
            const response = await fetch("/api/links", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({
                    "url": url,
                }),
            });
            if (response.ok) {
                this.hideErrText("");
                this.hideBtnContainer();
                this.hideBtnText();
                this.showResultContainer();
                this.setLink((await response.json())["short_url"]);
            } else {
                this.showErrText("На сервере произошла ошибка");
            }
        } catch {
            this.showErrText("Сервер не ответил");
        }
    };
}({
    "input": document.getElementById("inp"),
    "linkURL": document.getElementById("linkURL"),
    "linkText": document.getElementById("linkText"),
    "resultContainer": document.getElementsByClassName("shortener__result-container")[0],
    "btnContainer": document.getElementsByClassName("shortener__post-button-container")[0],
    "btnText": document.getElementsByClassName("shortener__post-button-text")[0],
    "errText": document.getElementsByClassName("shortener__error-text")[0],
});
