FROM debian:trixie-slim
RUN apt-get update && apt-get install -y ffmpeg curl && rm -rf /var/lib/apt/lists/*
WORKDIR /app
# Download the latest release from GitHub
RUN curl -s https://api.github.com/repos/diego09310/rockfm-scraper/releases/latest \
    | grep "browser_download_url" \
    | grep "rockfmScraper" \
    | cut -d '"' -f 4 \
    | xargs curl -L -o rockfmScraper \
    && chmod +x rockfmScraper

CMD ["./rockfmScraper", "--db"]
