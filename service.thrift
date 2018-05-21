namespace * price

service PriceService {
  string price(1:string key);
  list<string> prices(1:string key, 2:i16 start, 3:i16 stop);
}