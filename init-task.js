const htmlFetcher = require('./pkg/HtmlFetcher');
const HtmlNode = require('./pkg/HtmlNode');
const MarkdownBuilder = require('./pkg/MarkdownBuilder');
const HtmlToMarkdown = require('./pkg/HtmlToMarkdown');
const taskDir = require('./pkg/TaskDir');
const ARGUMENT_COUNT = 1;


if (process.argv.length !== ARGUMENT_COUNT + 2) {
    console.log(`Required ${ARGUMENT_COUNT} arguments`);
    process.exit();
}
const URL = process.argv[2];

start(URL);

async function start(url) {
    const html = await htmlFetcher.get(url, '[data-cy="question-title"]');
    const htmlNode = new HtmlNode(html);
    const markdownBuilder = new MarkdownBuilder();
    const htmlToMarkdown = new HtmlToMarkdown(htmlNode.getQuestionHtml());
    markdownBuilder.addTitle(htmlNode.getTitle(), htmlNode.getDifficulty());
    markdownBuilder.addQuestion(htmlToMarkdown.convert());
    taskDir.init(htmlNode.getTitle(), markdownBuilder.getMarkDown());
}