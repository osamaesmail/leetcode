class MarkdownBuilder {
    markdown = '';

    addLine(line) {
        this.markdown += `${line}\n`;
        return this;
    }

    addTitle(title, difficulty) {
        this.addLine(`# ${title} \`${difficulty}\``);
        return this;
    }

    addQuestion(question) {
        this.addLine(question);
        return this;
    }

    getMarkDown() {
        return this.markdown;
    }
}

module.exports = MarkdownBuilder