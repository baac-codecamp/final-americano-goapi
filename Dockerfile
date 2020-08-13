FROM scratch
ADD ./bin/check-salak-go-api /
CMD ["/check-salak-go-api"]
