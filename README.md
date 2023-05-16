# Quanta
Quanta is an algorithmic stock trading software that is accessible through the web. The software utilizes proprietary ML models to make accurate predictions about where the direction of the market is headed, and acts on them. There is a combination of deep-learning, statistical, and sentiment based modeling that goes into the predictions. The entire portfolio is managed using statistical and mathematical financial strategies to ensure the least amount of risk.

User's can view and track the software's trades in real-time, and will explore the notion of profit sharing in the future.

# Architecture
![Quanta drawio](https://github.com/ndavidson19/quanta/assets/59320455/55c9435a-31d6-4443-a805-bcfa6e429831)

## Kafka
Apache Kafka is used for real time streaming of stock and news data.

## Spark
Spark is used quickly calculate the features needed to do ML model inference (combination of TA, SA, and other metrics), as well as complicated financial derivates pricing calculations.

## Portfolio Management Microservice
more to come (and a name)

## Prediction Microservice
more to come (and a name)

## Trading Microservice
more to come (and a name)
