
get-dependencies:
	dep ensure
	mkdir -p menu
	rm -f menu/*
	wget -P menu/ https://cdn.shopify.com/s/files/1/1519/8146/products/151_BenR_170130.jpg
	wget -P menu/ https://www.organicfacts.net/wp-content/uploads/banana.jpg
	wget -P menu/ https://www.recipetineats.com/wp-content/uploads/2018/05/Chicken-Pad-Thai_0.jpg
	wget -P menu/ http://consorcioserrano.es/wp-content/uploads/2017/10/sabor-aroma-central.jpg

run:
	go run generate-menu.go