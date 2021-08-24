const cheerio = require('cheerio');

class HtmlNode {
    $
    constructor(html) {
        this.$ = cheerio.load(html);
    }

    getTitle() {
        return this.$('[data-cy="question-title"]').text();
    }

    getDifficulty() {
        return this.$('[diff]').text();
    }

    getQuestionHtml() {
        return this.$('.question-content__JfgR').html();
    }
}

module.exports = HtmlNode