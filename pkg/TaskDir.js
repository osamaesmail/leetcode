const fs = require('fs');

function init(title, markdown) {
    const dir = `./${title.toLowerCase().replace(/\s+/mg, '-').replace(/\./mg, '')}`;

    if (!fs.existsSync(dir)) {
        fs.mkdirSync(dir);
    }

    fs.writeFileSync(`${dir}/README.md`, markdown);
}

module.exports = {
    init,
}