tailwind-build:
	pnpm tailwindcss -i ./assets/css/styles.css -o ./static/css/styles.css --minify

tailwind-watch:
	pnpm tailwindcss -i ./assets/css/styles.css -o ./static/css/styles.css --minify --watch

templ-generate:
	templ generate

templ-watch:
	templ generate -watch -proxy=http://localhost:8081
