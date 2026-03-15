new class {
    constructor(data) {
        this.data = data;
        this.data.input.addEventListener("input", () => {
            if (this.data["input"].value.length > 0) {
                this.hideDescriptionTextContainer();
                this.showFooterTextContainer();
                this.hideErrorText();
            } else {
                this.showDescriptionTextContainer();
                this.hideFooterTextContainer();
            };
            this.hideSection_2Container();
        });
        this.data.postButton.addEventListener("click", () => {
            if (!this.data.input.value.trim().length === 0 || !this.data.input.value.includes(".")) {
                this.showErrorText("Некорректная ссылка");
                this.hideFooterTextContainer();
            } else {
                this.post(this.data.input.value);
            };
        });
        this.data.copyButton_1.addEventListener("click", () => {
            try {
                navigator.clipboard.writeText(this.data.linkURL.href.replace("xn--q1a", "с")).then(() => {
                    alert("Скопировано!");
                }).catch(error => {
                    alert(`Ошибка копирования: ${error}`)
                });
            } catch (error) {
                alert(`Ошибка копирования: ${error}`)
            };
        });
        this.hideErrorText();
        this.hideFooterTextContainer();
        this.hideSection_2Container();
        if (document.location.search.length > 0 && document.location.search[0] === "?") {
            var args = document.location.search.split("=")
            if (args.length === 2 && args[0].length === 2 && args[0][1] == "c" && args[1].length > 0) {
                this.get(args[1]);
            };
        };
    };
    showErrorText(text) {
        this.data.errorTextContainer.classList.add("default-show");
        this.data.errorTextContainer.classList.remove("default-hide");
        this.data.errorText.textContent = text;
    };
    hideErrorText() {
        this.data.errorTextContainer.classList.add("default-hide");
        this.data.errorTextContainer.classList.remove("default-show");
        this.data.errorText.textContent = "";
    };
    showDescriptionTextContainer() {
        this.data.descriptionTextContainer.classList.add("default-show");
        this.data.descriptionTextContainer.classList.remove("default-hide");
    };
    hideDescriptionTextContainer() {
        this.data.descriptionTextContainer.classList.add("default-hide");
        this.data.descriptionTextContainer.classList.remove("default-show");
    };
    showFooterTextContainer() {
        this.data.footerTextContainer.classList.add("default-show");
        this.data.footerTextContainer.classList.remove("default-hide");
    };
    hideFooterTextContainer() {
        this.data.footerTextContainer.classList.add("default-hide");
        this.data.footerTextContainer.classList.remove("default-show");
    };
    showSection_2Container() {
        this.data.section_2Container.classList.add("default-show");
        this.data.section_2Container.classList.remove("default-hide");
    };
    hideSection_2Container() {
        this.data.section_2Container.classList.add("default-hide");
        this.data.section_2Container.classList.remove("default-show");
    };
    setLink(url) {
        this.data.linkText.textContent = url.replace("https://", "");
        this.data.linkURL.href = url;
    };
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
                this.hideFooterTextContainer();
            }
        } catch {
            this.showErrorText("Сервер не ответил");
            this.hideFooterTextContainer();
        };
    };
    async get(short_code) {
        try {
            const response = await fetch(`/api/links/${short_code}`, {
                method: "GET",
            });
            if (response.ok) {
                this.hideDescriptionTextContainer();
                this.showSection_2Container();
                ((reponse_data) => {
                    this.data.input.value = reponse_data["original_url"];
                    this.setLink(reponse_data["short_url"]);
                })(await response.json());
            } else {
                document.location.href = "/";
            }
        } catch (error) {
            document.location.href = "/";
        };
    };
}({
    input: document.getElementById("inp"),
    postButton: document.getElementById("postBtn"),
    copyButton_1: document.getElementById("copyBtn_1"),
    errorText: document.getElementById("errorText"),
    errorTextContainer: document.getElementById("errorTextContainer"),
    descriptionTextContainer: document.getElementById("descriptionTextContainer"),
    footerTextContainer: document.getElementById("footerTextContainer"),
    section_2Container: document.getElementById("section_2Container"),
    linkURL: document.getElementById("linkURL"),
    linkText: document.getElementById("linkText"),
});
