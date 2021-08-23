const cheerio = require('cheerio');
const TurndownService = require('turndown');

class HtmlToMarkdown {
    $;
    turndownService = new TurndownService();

    constructor(html) {
        this.$ = cheerio.load(html)
        this._initExtraRules();
        this._generateHtml();
    }

    convert() {
        const markdown = this.turndownService.turndown(this.html);
        return markdown.replace(/\\\[/gm, '[').replace(/\\]/gm, ']');
    }

    _generateHtml() {
        this.$('pre').each((i, node) => {
            const $ = this.$(node);
            $.html($.text());
        });

        this.html = this.$.html();
    }

    _initExtraRules() {
        this.turndownService.addRule('pre', {
            filter: ['pre'],
            replacement: (content) => '\n```\n' + content + '\n```\n',
        });
    }
}

module.exports = HtmlToMarkdown;