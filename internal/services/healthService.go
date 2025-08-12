package services

type HealthCheck struct{}

func (s *HealthCheck) HealthCheckServices() (string, error) {
	return "testing data health", nil
}
