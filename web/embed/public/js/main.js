new class {
    constructor(data) {
        this.data = data;
        this.data.input.addEventListener("input", () => {
            if (this.data["input"].value.length > 0) {
                this.hideDescriptionText();
                this.showFooterText();
                this.hideErrorText();
            } else {
                this.showDescriptionText();
                this.hideFooterText();
            };
            this.hideResultBlock();
            this.hideResultDescriptionText();
            this.hideResultTitleText();
        });
        this.data.postButton.addEventListener("click", () => {
            if (!this.data.input.value.trim().length === 0 || !this.data.input.value.includes(".")) {
                this.showErrorText("Некорректная ссылка");
            } else {
                this.post(this.data.input.value);
            }
        });
        this.data.copyButton.addEventListener("click", () => {
            try {
                navigator.clipboard.writeText(this.data.linkURL.href.replace("xn--q1a", "с")).then(() => {
                    alert("Скопировано!");
                }).catch(error => {
                    alert(`Ошибка копирования: ${error}`)
                });
            } catch (error) {
                alert(`Ошибка копирования: ${error}`)
            }
        });
        this.hideErrorText();
        this.hideFooterText();
        this.hideResultTitleText();
        this.hideResultDescriptionText();
        this.hideResultBlock();
        if (document.location.search.length > 0 && document.location.search[0] === "?") {
            var args = document.location.search.split("=")
            if (args.length === 2 && args[0].length === 2 && args[0][1] == "c" && args[1].length > 0) {
                this.get(args[1]);
            }
        }
    };
    showErrorText(text) {
        this.data.errorText.classList.add("shortener__error-text-show");
        this.data.errorText.classList.remove("shortener__error-text-hide");
        this.data.errorText.textContent = text;
        this.hideFooterText();
    };
    hideErrorText() {
        this.data.errorText.classList.add("shortener__error-text-hide");
        this.data.errorText.classList.remove("shortener__error-text-show");
        this.data.errorText.textContent = "";
    };
    showFooterText() {
        this.data.footerText.classList.add("shortener__footer-text-show");
        this.data.footerText.classList.remove("shortener__footer-text-hide");
    };
    hideFooterText() {
        this.data.footerText.classList.add("shortener__footer-text-hide");
        this.data.footerText.classList.remove("shortener__footer-text-show");
    };
    showDescriptionText() {
        this.data.descriptionText.classList.add("shortener__description-text-show");
        this.data.descriptionText.classList.remove("shortener__description-text-hide");
    };
    hideDescriptionText() {
        this.data.descriptionText.classList.add("shortener__description-text-hide");
        this.data.descriptionText.classList.remove("shortener__description-text-show");
    };
    showResultTitleText() {
        this.data.resultTitleText.classList.add("shortener__result-title-text-show");
        this.data.resultTitleText.classList.remove("shortener__result-title-text-hide");
    };
    hideResultTitleText() {
        this.data.resultTitleText.classList.add("shortener__result-title-text-hide");
        this.data.resultTitleText.classList.remove("shortener__result-title-text-show");
    };
    showResultDescriptionText() {
        this.data.resultDescriptionText.classList.add("shortener__result-description-text-show");
        this.data.resultDescriptionText.classList.remove("shortener__result-description-text-hide");
    };
    hideResultDescriptionText() {
        this.data.resultDescriptionText.classList.add("shortener__result-description-text-hide");
        this.data.resultDescriptionText.classList.remove("shortener__result-description-text-show");
    };
    showResultBlock() {
        this.data.resultBlock.classList.add("shortener__result-block-show");
        this.data.resultBlock.classList.remove("shortener__result-block-hide");
    };
    hideResultBlock() {
        this.data.resultBlock.classList.add("shortener__result-block-hide");
        this.data.resultBlock.classList.remove("shortener__result-block-show");
    };
    setLink(url) {
        this.data.linkText.textContent = url.replace("https://", "");
        this.data.linkURL.href = url;
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
                document.location.href = `/?c=${(await response.json())["short_code"]}`;
            } else {
                this.showErrorText(`Ошибка с сервера: ${await response.text()}`);
            }
        } catch {
            this.showErrorText("Сервер не ответил");
        };
    };
    async get(short_code) {
        try {
            const response = await fetch(`/api/links/${short_code}`, {
                method: "GET",
            });
            if (response.ok) {
                this.hideDescriptionText();
                this.showResultTitleText();
                this.showResultDescriptionText();
                this.showResultBlock();
                ((reponse_data) => {
                    this.data.input.value = reponse_data["original_url"];
                    this.setLink(reponse_data["short_url"]);
                })(await response.json());
            } else {
                document.location.href = "/";
            }
        } catch {
            document.location.href = "/";
        };
    }
}({
    input: document.getElementById("inp"),
    linkURL: document.getElementById("linkURL"),
    linkText: document.getElementById("linkText"),
    postButton: document.getElementById("postBtn"),
    copyButton: document.getElementById("copyBtn"),
    errorText: document.getElementsByClassName("shortener__error-text")[0],
    footerText: document.getElementsByClassName("shortener__footer-text")[0],
    descriptionText: document.getElementsByClassName("shortener__description-text")[0],
    resultTitleText: document.getElementsByClassName("shortener__result-title-text")[0],
    resultDescriptionText: document.getElementsByClassName("shortener__result-description-text")[0],
    resultBlock: document.getElementsByClassName("shortener__result-block")[0],
});
