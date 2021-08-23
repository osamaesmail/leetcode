const puppeteer = require('puppeteer');

async function get(url, selector) {
    const browser = await puppeteer.launch();
    const page = await browser.newPage();
    await page.goto(url);
    await page.waitForSelector(selector);
    const html = await page.evaluate(() => document.querySelector('*').outerHTML);
    await browser.close();
    return html;
}

module.exports = {
    get,
}