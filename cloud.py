/*
1, SQS, S3, EC2
2,  SimpleDB, Mechanical Turk, Elastic Block Store, Elastic Beanstalk, Relational Database Service, DynamoDB, 
      CloudWatch, Simple Workflow, CloudFront, and Availability Zones


Amazon Mechanical Turk (MTurk) is a crowdsourcing website for businesses to hire remotely located "crowdworkers" to 
      perform discrete on-demand tasks that computers are currently unable to do. It is operated under Amazon Web Services, 
            and is owned by Amazon. Employers (known as requesters) post jobs known as Human Intelligence Tasks (HITs), 
                  such as identifying specific content in an image or video, writing product descriptions, or answering survey questions. 
                        Workers, colloquially known as Turkers or crowdworkers, browse among existing jobs and complete them in exchange for 
                              a fee set by the employer. To place jobs, the requesting programs use an open application programming interface (API), 
                                    or the more limited MTurk Requester site. As of April 2019, Requesters could register from 49 approved countries.


AWS Elastic Beanstalk is an orchestration service offered by Amazon Web Services for deploying applications which orchestrates various AWS services, 
      including EC2, S3, Simple Notification Service (SNS), CloudWatch, autoscaling, and Elastic Load Balancers. Elastic Beanstalk provides an 
            additional layer of abstraction over the bare server and OS; users instead see a pre-built combination of OS and platform, 
                   such as "64bit Amazon Linux 2014.03 v1.1.0 running Ruby 2.0 (Puma)" or "64bit Debian jessie v2.0.7 running Python 3.4 (Preconfigured - Docker)".
                        Deployment requires a number of components to be defined: an 'application' as a logical container for the project, a 'version' 
                              which is a deployable build of the application executable, a 'configuration template' that contains configuration information for both 
                                    the Beanstalk environment and for the product. Finally an 'environment' combines a 'version' with a 'configuration' and 
                                          deploys them. Executables themselves are uploaded as archive files to S3 beforehand and the 'version' is just a pointer to this.

The Amazon Simple Workflow Service (Amazon SWF) makes it easy to build applications that coordinate work across distributed components. In Amazon SWF, 
      a task represents a logical unit of work that is performed by a component of your application. Coordinating tasks across the application involves 
            managing intertask dependencies, scheduling, and concurrency in accordance with the logical flow of the application. Amazon SWF gives you full control 
                    over implementing tasks and coordinating them without worrying about underlying complexities such as tracking their progress and maintaining 
                        their state.
                                          








*/
