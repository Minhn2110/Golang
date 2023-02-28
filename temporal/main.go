package main

import (
	"context"
	"fmt"

	"go.temporal.io/api/common/v1"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/api/history/v1"
	temporal "go.temporal.io/sdk/client"
	"go.temporal.io/sdk/converter"
)

type WorkflowExecutionHistoryItemV2 struct {
	EventId      int64
	HistoryItems []interface{}
}

func main() {
	ctx := context.Background()
	temporalClient, _ := temporal.Dial(temporal.Options{
		HostPort: "127.0.0.1:8082",
	})
	var historyItems []interface{}
	// var workflowDSL interface{}
	// var items []interface{}

	iter := temporalClient.GetWorkflowHistory(ctx, "4b29e3cd-3baf-4248-917f-301634e55c45", "", false, 1)
	for iter.HasNext() {
		event, _ := iter.Next()
		m, _ := DecodeWorkflowExecutionLogEvent(event)
		fmt.Println("m", m)
		if m != nil {
			historyItems = append(historyItems, m)
		}

		// if IsWorkflowDSL(event) {
		// 	workflowDSL, err := DecodeWorkflowDSL(event)
		// }

		// if IsLogEvent(event) {
		// 	item, err := DecodeLogEvent(event)
		// 	if err != nil {
		// 		return nil, err
		// 	}

		// 	items = append(items, item)
		// }
	}

	// flat := lo.Flatten[int]([][]int{{0, 1}, {2, 3, 4, 5}})

	// fmt.Println("historyItems", lo.Flatten([][]any{historyItems}))
	fmt.Println("historyItems", historyItems)
}

func DecodePayload(payload *common.Payload) (interface{}, error) {
	conv := converter.GetDefaultDataConverter()

	var data interface{}
	err := conv.FromPayload(payload, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode payload: %w", err)
	}

	return data, nil
}

func DecodePayloads(payloads *common.Payloads) (map[string]any, error) {
	conv := converter.GetDefaultDataConverter()

	var data map[string]any
	err := conv.FromPayloads(payloads, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode payload: %w", err)
	}

	return data, nil
}

func DecodeWorkflowExecutionLogEvent(event *history.HistoryEvent) (interface{}, error) {

	// status := "running"
	// switch {
	// case event.GetEventType() == enums.EVENT_TYPE_WORKFLOW_EXECUTION_COMPLETED:
	// 	status = "Completed"
	// case event.GetEventType() == enums.EVENT_TYPE_WORKFLOW_EXECUTION_FAILED:
	// 	status = "Failed"
	// }

	// fmt.Println("status", status)

	// var dsl interface{}

	// if event.EventType == enums.EVENT_TYPE_WORKFLOW_EXECUTION_STARTED {
	// 	attrs := event.GetWorkflowExecutionStartedEventAttributes()
	// 	payloads := attrs.GetInput().GetPayloads()
	// 	payloadCount := len(payloads)
	// 	if payloadCount != 1 {
	// 		return nil, fmt.Errorf("invalid payload count: %d", payloadCount)
	// 	}
	// 	input, _ := DecodePayload(payloads[0])
	// 	dsl = input
	// 	fmt.Println("dsl", dsl.(map[string]interface{})["original_workflow"])
	// 	// fmt.Printf("type of a is %T\n", dsl)
	// }

	if event.GetEventType() == enums.EVENT_TYPE_MARKER_RECORDED {
		details := event.GetMarkerRecordedEventAttributes().GetDetails()
		markerName := event.GetMarkerRecordedEventAttributes().MarkerName

		if markerName == "LocalActivity" {
			data, _ := DecodePayloads(details["data"])
			if data["ActivityType"] == "log_event" {
				result, _ := DecodePayloads(details["result"])
				result["time"] = event.EventTime
				fmt.Println("result", result)
				return result, nil
			}
		}
		return nil, nil
	}

	return nil, nil
}
