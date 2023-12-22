module.exports = {
    content: [
        './internal/templates/**/*{_templ.go,.templ}',
    ],
    plugins: [
        require('@tailwindcss/forms'),
        require('@tailwindcss/typography'),
    ]
}