serve:
	go run cmd/server/main.go
css:
	npx tailwindcss -i ./web/index.css -o ./web/static/styles/index.css --watch
templ:
	watchexec -e templ -r templ generate
