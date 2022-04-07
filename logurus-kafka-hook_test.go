package logurus_kafka_hook

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"reflect"
	"testing"
	"time"
)

func TestKafkaHook_Fire(t *testing.T) {
	type fields struct {
		id        string
		levels    []logrus.Level
		formatter logrus.Formatter
		producer  sarama.AsyncProducer
	}
	type args struct {
		entry *logrus.Entry
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hook := &KafkaHook{
				id:        tt.fields.id,
				levels:    tt.fields.levels,
				formatter: tt.fields.formatter,
				producer:  tt.fields.producer,
			}
			if err := hook.Fire(tt.args.entry); (err != nil) != tt.wantErr {
				t.Errorf("Fire() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewKafkaHook(t *testing.T) {
	type args struct {
		id        string
		levels    []logrus.Level
		formatter logrus.Formatter
		brokers   []string
	}
	tests := []struct {
		name    string
		args    args
		want    *KafkaHook
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewKafkaHook(tt.args.id, tt.args.levels, tt.args.formatter, tt.args.brokers)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewKafkaHook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKafkaHook() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKafkaHook(t *testing.T) {
	// Create a new KafkaHook
	hook, err := NewKafkaHook(
		"kh",
		[]logrus.Level{logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel},
		&logrus.JSONFormatter{},
		[]string{"192.168.60.5:9092", "192.168.60.6:9092", "192.168.60.7:9092"},
	)

	if err != nil {
		t.Errorf("Can not create KafkaHook: %v\n", err)
	}

	// Create a new logrus.Logger
	logger := logrus.New()

	// Add hook to logger
	logger.Hooks.Add(hook)

	t.Logf("logger: %v", logger)
	t.Logf("logger.Out: %v", logger.Out)
	t.Logf("logger.Formatter: %v", logger.Formatter)
	t.Logf("logger.Hooks: %v", logger.Hooks)
	t.Logf("logger.Level: %v", logger.Level)

	// Add topics
	l := logger.WithField("topics", []string{"topic_1", "topic_2", "topic_3"})

	l.Debug("This must not be logged")

	l.Info("This is an Info msg")

	l.Warn("This is a Warn msg")

	l.Error("This is an Error msg")

	// Ensure log messages were written to Kafka
	time.Sleep(time.Second)
}
