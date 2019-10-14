build:
	rm nc-test-app/dist -rf
	cd frontend && npm run build
	cp frontend/dist nc-test-webapp -r
