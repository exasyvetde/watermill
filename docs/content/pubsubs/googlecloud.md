+++
title = "Google Cloud Pub/Sub"
description = "The fully-managed real-time messaging service from Google"
date = 2019-07-06T22:30:00+02:00
bref = "The fully-managed real-time messaging service from Google"
weight = -60
type = "docs"
toc = false
+++

### Google Cloud Pub/Sub

Cloud Pub/Sub brings the flexibility and reliability of enterprise message-oriented middleware to
the cloud.

At the same time, Cloud Pub/Sub is a scalable, durable event ingestion and delivery
system that serves as a foundation for modern stream analytics pipelines.
By providing many-to-many, asynchronous messaging that decouples senders and receivers,
it allows for secure and highly available communication among independently written applications.

Cloud Pub/Sub delivers low-latency, durable messaging that helps developers quickly integrate
systems hosted on the Google Cloud Platform and externally.

Documentation: [https://cloud.google.com/pubsub/docs/](https://cloud.google.com/pubsub/docs/overview)

#### Characteristics

| Feature | Implements | Note |
| ------- | ---------- | ---- |
| ConsumerGroups | yes | multiple subscribers within the same Subscription name  |
| ExactlyOnceDelivery | no |  |
| GuaranteedOrder | no | |
| Persistent | yes* | maximum retention time is 7 days |

#### Configuration

{{% render-md %}}
{{% load-snippet-partial file="src-link/watermill-googlecloud/pkg/googlecloud/publisher.go" first_line_contains="type PublisherConfig struct " last_line_contains="func NewPublisher" %}}
{{% /render-md %}}

{{% render-md %}}
{{% load-snippet-partial file="src-link/watermill-googlecloud/pkg/googlecloud/subscriber.go" first_line_contains="type SubscriberConfig struct {" last_line_contains="func NewSubscriber(" %}}
{{% /render-md %}}

##### Subscription name

To receive messages published to a topic, you must create a subscription to that topic.
Only messages published to the topic after the subscription is created are available to subscriber
applications.

The subscription connects the topic to a subscriber application that receives and processes
messages published to the topic.

A topic can have multiple subscriptions, but a given subscription belongs to a single topic.

In Watermill, the subscription is created automatically during calling `Subscribe()`.
Subscription name is generated by function passed to `SubscriberConfig.GenerateSubscriptionName`.
By default, it is just the topic name (`TopicSubscriptionName`).

When you want to consume messages from a topic with multiple subscribers, you should use
`TopicSubscriptionNameWithSuffix` or your custom function to generate the subscription name.

#### Connecting

Watermill will connect to the instance of Google Cloud Pub/Sub indicated by the environment variables. For production setup, set the `GOOGLE_APPLICATION_CREDENTIALS` env, as described in [the official Google Cloud Pub/Sub docs](https://cloud.google.com/pubsub/docs/quickstart-client-libraries#pubsub-client-libraries-go). Note that you won't need to install the Cloud SDK, as Watermill will take care of the administrative tasks (creating topics/subscriptions) with the default settings and proper permissions.

For development, you can use a Docker image with the emulator and the `PUBSUB_EMULATOR_HOST` env ([check out the Getting Started guide]({{< ref "getting-started#subscribing_gcloud" >}})).

{{% render-md %}}
{{% load-snippet-partial file="src-link/_examples/pubsubs/googlecloud/main.go" first_line_contains="publisher, err :=" last_line_contains="panic(err)" padding_after="1" %}}
{{% /render-md %}}

{{% render-md %}}
{{% load-snippet-partial file="src-link/_examples/pubsubs/googlecloud/main.go" first_line_contains="subscriber, err :=" last_line_contains="panic(err)" padding_after="1" %}}
{{% /render-md %}}

#### Publishing

{{% render-md %}}
{{% load-snippet-partial file="src-link/watermill-googlecloud/pkg/googlecloud/publisher.go" first_line_contains="// Publish" last_line_contains="func (p *Publisher) Publish" %}}
{{% /render-md %}}

#### Subscribing

{{% render-md %}}
{{% load-snippet-partial file="src-link/watermill-googlecloud/pkg/googlecloud/subscriber.go" first_line_contains="// Subscribe " last_line_contains="func (s *Subscriber) Subscribe" %}}
{{% /render-md %}}

#### Marshaler

Watermill's messages cannot be directly sent to Google Cloud Pub/Sub - they need to be marshaled. You can implement your marshaler or use the default implementation.

{{% render-md %}}
{{% load-snippet-partial file="src-link/watermill-googlecloud/pkg/googlecloud/marshaler.go" first_line_contains="// Marshaler" last_line_contains="type DefaultMarshalerUnmarshaler " padding_after="0" %}}
{{% /render-md %}}

