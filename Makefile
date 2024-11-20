# Main Docker Commands
# ======================

build:
	@sudo docker build -t ses_back .

run:
	@sudo docker run -d --name ses_back -p 8080:8080 ses_back