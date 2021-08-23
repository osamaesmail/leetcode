const fs = require('fs');

function init(title, markdown) {
    const dir = `./${title.toLowerCase().replace(/\s+/, '-').replace(/\./, '')}`;

    if (!fs.existsSync(dir)) {
        fs.mkdirSync(dir);
    }

    fs.writeFileSync(`${dir}/README.md`, markdown);
}

module.exports = {
    init,
}