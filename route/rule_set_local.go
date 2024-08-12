package route

import (
	"context"

	"github.com/sagernet/sing-box/adapter"
	C "github.com/sagernet/sing-box/constant"
	"github.com/sagernet/sing-box/option"
)

var _ adapter.RuleSet = (*LocalRuleSet)(nil)

type LocalRuleSet struct {
	abstractRuleSet
}

func NewLocalRuleSet(ctx context.Context, router adapter.Router, options option.RuleSet) (*LocalRuleSet, error) {
	ruleSet := LocalRuleSet{
		abstractRuleSet: abstractRuleSet{
			ctx:    ctx,
			tag:    options.Tag,
			rType:  C.TypeLocal,
			path:   options.Path,
			format: options.Format,
		},
	}
	return &ruleSet, ruleSet.loadFromFile(router, true)
}

func (s *LocalRuleSet) Update(router adapter.Router) error {
	return s.loadFromFile(router, false)
}

func (s *LocalRuleSet) StartContext(_ context.Context, _ adapter.RuleSetStartContext) error {
	return nil
}

func (s *LocalRuleSet) Close() error {
	return nil
}
