package factory

import "fmt"

func (p *ClientProvider) Close() error {
	p.clientsMutex.Lock()
	defer p.clientsMutex.Unlock()

	var errs []error
	for st, client := range p.clients {
		if closer, ok := client.(interface{ Close() error }); ok {
			if err := closer.Close(); err != nil {
				errs = append(errs, fmt.Errorf("%s: %w", st, err))
			}
		}
		delete(p.clients, st)
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors closing clients: %v", errs)
	}
	return nil
}
