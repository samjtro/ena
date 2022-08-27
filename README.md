# stock-news

super fast stock news lookup from the command line.

## usage

```
0. go install github.com/samjtro/sn@latest
1. sn --t ["multi-keyword" or "mkw", "keyword" or "kw"] --kw [keyword(s) of your choice seperated by ","] --s ["prnewswire" or "prn", "google" or "g"]

example return:

$ sn --t mkw --kw biocardia,bcda

Headline:  Has BioCardia (BCDA) Outpaced Other Medical Stocks This Year?
URL:  https://news.google.com/articles/CBMiWmh0dHBzOi8vd3d3Lm5hc2RhcS5jb20vYXJ0aWNsZXMvaGFzLWJpb2NhcmRpYS1iY2RhLW91dHBhY2VkLW90aGVyLW1lZGljYWwtc3RvY2tzLXRoaXMteWVhctIBXmh0dHBzOi8vd3d3Lm5hc2RhcS5jb20vYXJ0aWNsZXMvaGFzLWJpb2NhcmRpYS1iY2RhLW91dHBhY2VkLW90aGVyLW1lZGljYWwtc3RvY2tzLXRoaXMteWVhcj9hbXA?hl=en-US&gl=US&ceid=US%3Aen

Headline:  BioCardia Reports Second Quarter 2022 Business Highlights And Financial Results
URL:  https://news.google.com/articles/CBMiU2h0dHBzOi8vZmluYW5jZS55YWhvby5jb20vbmV3cy9iaW9jYXJkaWEtcmVwb3J0cy1zZWNvbmQtcXVhcnRlci0yMDIyLTIwMDUwMDc5OC5odG1s0gFbaHR0cHM6Ly9maW5hbmNlLnlhaG9vLmNvbS9hbXBodG1sL25ld3MvYmlvY2FyZGlhLXJlcG9ydHMtc2Vjb25kLXF1YXJ0ZXItMjAyMi0yMDA1MDA3OTguaHRtbA?hl=en-US&gl=US&ceid=US%3Aen

Headline:  Will BioCardia (NASDAQ:BCDA) Spend Its Cash Wisely?
URL:  https://news.google.com/articles/CBMiggFodHRwczovL3NpbXBseXdhbGwuc3Qvc3RvY2tzL3VzL3BoYXJtYWNldXRpY2Fscy1iaW90ZWNoL25hc2RhcS1iY2RhL2Jpb2NhcmRpYS9uZXdzL3dpbGwtYmlvY2FyZGlhLW5hc2RhcWJjZGEtc3BlbmQtaXRzLWNhc2gtd2lzZWx50gGGAWh0dHBzOi8vc2ltcGx5d2FsbC5zdC9zdG9ja3MvdXMvcGhhcm1hY2V1dGljYWxzLWJpb3RlY2gvbmFzZGFxLWJjZGEvYmlvY2FyZGlhL25ld3Mvd2lsbC1iaW9jYXJkaWEtbmFzZGFxYmNkYS1zcGVuZC1pdHMtY2FzaC13aXNlbHkvYW1w?hl=en-US&gl=US&ceid=US%3Aen

Headline:  BioCardia Inc (BCDA) has gained 8.33% in a Week, Should You Sell?
URL:  https://news.google.com/articles/CBMicGh0dHBzOi8vd3d3LmludmVzdG9yc29ic2VydmVyLmNvbS9uZXdzL3N0b2NrLXVwZGF0ZS9iaW9jYXJkaWEtaW5jLWJjZGEtaGFzLWdhaW5lZC04LTMzLWluLWEtd2Vlay1zaG91bGQteW91LXNlbGzSAXRodHRwczovL3d3dy5pbnZlc3RvcnNvYnNlcnZlci5jb20vbmV3cy9zdG9jay11cGRhdGUvYW1wL2Jpb2NhcmRpYS1pbmMtYmNkYS1oYXMtZ2FpbmVkLTgtMzMtaW4tYS13ZWVrLXNob3VsZC15b3Utc2VsbA?hl=en-US&gl=US&ceid=US%3Aen

Headline:  BIOCARDIA, INC. : Regulation FD Disclosure, Financial Statements and Exhibits (form 8-K)
URL:  https://news.google.com/articles/CBMiogFodHRwczovL3d3dy5tYXJrZXRzY3JlZW5lci5jb20vcXVvdGUvc3RvY2svQklPQ0FSRElBLUlOQy02NDMwMjMzMC9uZXdzL0JJT0NBUkRJQS1JTkMtUmVndWxhdGlvbi1GRC1EaXNjbG9zdXJlLUZpbmFuY2lhbC1TdGF0ZW1lbnRzLWFuZC1FeGhpYml0cy1mb3JtLTgtSy00MTM5ODk2MS_SAaYBaHR0cHM6Ly93d3cubWFya2V0c2NyZWVuZXIuY29tL2FtcC9xdW90ZS9zdG9jay9CSU9DQVJESUEtSU5DLTY0MzAyMzMwL25ld3MvQklPQ0FSRElBLUlOQy1SZWd1bGF0aW9uLUZELURpc2Nsb3N1cmUtRmluYW5jaWFsLVN0YXRlbWVudHMtYW5kLUV4aGliaXRzLWZvcm0tOC1LLTQxMzk4OTYxLw?hl=en-US&gl=US&ceid=US%3Aen

Headline:  BioCardia Announces First Canadian Clinical Site for CardiAMP Cell Therapy Heart Failure Trial
URL:  https://news.google.com/articles/CBMimQFodHRwczovL3d3dy5zdHJlZXRpbnNpZGVyLmNvbS9HbG9iZStOZXdzd2lyZS9CaW9DYXJkaWErQW5ub3VuY2VzK0ZpcnN0K0NhbmFkaWFuK0NsaW5pY2FsK1NpdGUrZm9yK0NhcmRpQU1QK0NlbGwrVGhlcmFweStIZWFydCtGYWlsdXJlK1RyaWFsLzIwMzkzMTc2Lmh0bWzSAQA?hl=en-US&gl=US&ceid=US%3Aen

Headline:  BIOCARDIA, INC. MANAGEMENT'S DISCUSSION AND ANALYSIS OF FINANCIAL CONDITION AND RESULTS OF OPERATIONS (form 10-Q)
URL:  https://news.google.com/articles/CBMisgFodHRwczovL3d3dy5tYXJrZXRzY3JlZW5lci5jb20vcXVvdGUvc3RvY2svQklPQ0FSRElBLUlOQy02NDMwMjMzMC9uZXdzL0JJT0NBUkRJQS1JTkMtTUFOQUdFTUVOVC1TLURJU0NVU1NJT04tQU5ELUFOQUxZU0lTLU9GLUZJTkFOQ0lBTC1DT05ESVRJT04tQU5ELVJFU1VMVFMtT0YtT1BFUkFUSU8tNDEyNzEwOTYv0gG2AWh0dHBzOi8vd3d3Lm1hcmtldHNjcmVlbmVyLmNvbS9hbXAvcXVvdGUvc3RvY2svQklPQ0FSRElBLUlOQy02NDMwMjMzMC9uZXdzL0JJT0NBUkRJQS1JTkMtTUFOQUdFTUVOVC1TLURJU0NVU1NJT04tQU5ELUFOQUxZU0lTLU9GLUZJTkFOQ0lBTC1DT05ESVRJT04tQU5ELVJFU1VMVFMtT0YtT1BFUkFUSU8tNDEyNzEwOTYv?hl=en-US&gl=US&ceid=US%3Aen

Headline:  Recap: BioCardia Q2 Earnings
URL:  https://news.google.com/articles/CBMiUWh0dHBzOi8vd3d3LmJlbnppbmdhLmNvbS9uZXdzL2Vhcm5pbmdzLzIyLzA4LzI4NDQ1MzkxL3JlY2FwLWJpb2NhcmRpYS1xMi1lYXJuaW5nc9IBLWh0dHBzOi8vd3d3LmJlbnppbmdhLmNvbS9hbXAvY29udGVudC8yODQ0NTM5MQ?hl=en-US&gl=US&ceid=US%3Aen

...

```