Json sended via broker-microservice to end point ---> broker-microservice /handdleSubmission/log

RabbitMq call pushqueue function :

setup the emiiter(NewEventEmiiter)-->declare exchange(stup function) --> setup the Emmiter -->publishing to rabbitmq to with (Push func) -->
-> //handler
 setup EventEmiiter & func NewEventEmmiter for getting emmiter  --> emmiter.push() --> pushToQueue -> return acknolegment json from that gateway microservice(broker)