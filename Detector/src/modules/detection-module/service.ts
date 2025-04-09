import { DetectionRequest, DetectionResponse } from './dtos'

export class DetectionService {
    /**
     * Main detection method
     */
    public static async detect(request: DetectionRequest): Promise<DetectionResponse> {
        const detectionResult = await this.makeApiCall(request)

        return new DetectionResponse({
            request,
            detectionInfo: {
                detected: detectionResult,
            },
        })
    }

    /**
     * Makes a GET request using the wallet address as a query param
     * Returns true if suspicious, false if API confirms clean (200 status)
     */
    public static async makeApiCall(request: DetectionRequest): Promise<boolean> {
        let detectionResult = true

        try {
            const walletAddress = request.trace.from
            const query = new URLSearchParams({ wallet: walletAddress }).toString()

            const response = await fetch(`https://your-api-endpoint.com/check?${query}`, {
                method: 'GET',
            })

            if (response.status === 200) {
                detectionResult = false
            }
        } catch (error) {
            console.error('API call failed:', error)
        }

        return detectionResult
    }
}
