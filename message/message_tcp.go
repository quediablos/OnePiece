package message

import "fmt"

const tcpFrameVersion = "1"

func GenerateAcquireLockTcpResponse(resourceId string) string {
	return fmt.Sprintf("%s|RES|SUCCESSFUL|LOCK|%s\n", tcpFrameVersion, resourceId)
}

func GenerateAcquireLockTcpFailedResponse(resourceId, errorCode, errorMessage string) string {
	return fmt.Sprintf("%s|RES|FAILED|LOCK|%s|%s|%s\n", tcpFrameVersion, resourceId, errorCode, errorMessage)
}

func GenerateReleaseLockTcpResponse(resourceId string) string {
	return fmt.Sprintf("%s|RES|SUCCESS|UNLOCK|%s\n", tcpFrameVersion, resourceId)
}

func GenerateReleaseLockTcpFailedResponse(resourceId, errorCode, errorMessage string) string {
	return fmt.Sprintf("%s|RES|FAILED|UNLOCK|%s|%s|%s\n", tcpFrameVersion, resourceId, errorCode, errorMessage)
}

// Stock responses

func GenerateReserveStockSuccessfulTcpResponse(reserveId string, resourceId string) string {
	// 1|RES|SUCCESSFUL|RESERVE_STOCK|<resourceId>|<reserveId>
	return fmt.Sprintf("%s|RES|SUCCESSFUL|RESERVE_STOCK|%s|%s\n", tcpFrameVersion, resourceId, reserveId)
}

func GenerateReserveStockFailedTcpResponse(resourceId string) string {
	// 1|RES|FAILED|RESERVE_STOCK|<resourceId>
	return fmt.Sprintf("%s|RES|FAILED|RESERVE_STOCK|%s\n", tcpFrameVersion, resourceId)
}

func GenerateCreateStockSuccessfulTcpResponse(resourceId string, quantity string) string {
	// 1|RES|SUCCESSFUL|CREATE_STOCK|<resourceId>|<quantity>
	return fmt.Sprintf("%s|RES|SUCCESSFUL|CREATE_STOCK|%s|%s\n", tcpFrameVersion, resourceId, quantity)
}

func GenerateCreateStockFailedTcpResponse(resourceId string) string {
	return fmt.Sprintf("%s|RES|FAILED|CREATE_STOCK|%s|STOCK_ALREADY_CREATED|Stock already created.\n",
		tcpFrameVersion, resourceId)
}

func GenerateReleaseStockSuccessfulTcpResponse(resourceId string) string {
	return fmt.Sprintf("%s|RES|SUCCESSFUL|RELEASE_STOCK|%s\n", tcpFrameVersion, resourceId)
}

func GenerateWaitRlSuccessfulTcpResponse() string {
	return fmt.Sprintf("%s|RES|SUCCESSFUL|RL_WAIT|%s\n", tcpFrameVersion, false)
}

func GenerateWaitRlFailedTcpResponse(errorCode string, errorMessage string) string {
	return fmt.Sprintf("%s|RES|FAILED|RL_WAIT|%s|%s\n", tcpFrameVersion, errorCode, errorMessage)
}

func GenerateSetupRlSuccessfulTcpResponse() string {
	return fmt.Sprintf("%s|RES|SUCCESS|RL_SETUP\n", tcpFrameVersion)
}

// Error responses

func GenerateGenericErrorTcpResponse() string {
	return fmt.Sprintf("%s|RES|FAILED|ERROR\n", tcpFrameVersion)
}

func GenerateErrorTcpResponse(err string) string {
	return fmt.Sprintf("%s|RES|FAILED|ERROR|%s\n", tcpFrameVersion, err)
}
